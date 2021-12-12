package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/kumparan-article-service/pkg/model"
	"github.com/ssentinull/kumparan-article-service/pkg/utils"
)

type articleRepository struct {
	cacher model.Cacher
	db     *sql.DB
}

func NewArticleRepository(db *sql.DB, c model.Cacher) model.ArticleRepository {
	return &articleRepository{
		cacher: c,
		db:     db,
	}
}

func (ar *articleRepository) CalculateVectors(ctx context.Context, article *model.Article) error {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"article": utils.Dump(article),
	})

	query := `UPDATE articles SET title_body_vectors = to_tsvector($1 || ' ' || $2),
		author_vectors = to_tsvector($3) WHERE id = $4`
	stmt, err := ar.db.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)

		return err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	_, err = stmt.ExecContext(ctx, article.Title, article.Body, article.Author, article.ID)
	if err != nil {
		logger.Error(err)
		err = model.ErrBadRequest

		return err
	}

	return nil
}

func (ar *articleRepository) Create(ctx context.Context, article *model.Article) error {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"article": utils.Dump(article),
	})

	query := `INSERT INTO articles (author, title, body, created_at) 
		VALUES ($1, $2, $3, $4) RETURNING id`
	stmt, err := ar.db.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)

		return err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	jkt, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		logger.Error(err)

		return err
	}

	lastInsertID := int64(0)
	now := time.Now().In(jkt)
	err = stmt.QueryRowContext(ctx, article.Author, article.Title, article.Body, now).Scan(&lastInsertID)
	if err != nil {
		logger.Error(err)
		err = model.ErrBadRequest

		return err
	}

	article.ID = lastInsertID
	article.CreatedAt = now

	return nil
}

func (ar *articleRepository) Read(ctx context.Context, qp model.QueryBuilder) ([]model.Article, error) {
	logger := logrus.WithFields(logrus.Fields{
		"context":    utils.Dump(ctx),
		"queryParam": qp,
	})

	baseQuery := "SELECT id, author, title, body, created_at FROM articles "
	orderQuery := "ORDER BY created_at DESC "
	filterQuery := qp.BuildWhereClause()
	query := baseQuery + filterQuery + orderQuery

	reply, err := ar.cacher.Get(query)
	if err != nil && err != bigcache.ErrEntryNotFound {
		logger.Error(err)

		return nil, err
	}

	if reply != nil {
		articles := make([]model.Article, 0)
		if err = json.Unmarshal(reply, &articles); err != nil {
			return nil, err
		}

		return articles, nil
	}

	rows, err := ar.db.QueryContext(ctx, query)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	articles := make([]model.Article, 0)
	for rows.Next() {
		var ID sql.NullInt64
		var author, title, body sql.NullString
		var createdAt sql.NullTime
		err = rows.Scan(&ID, &author, &title, &body, &createdAt)
		if err != nil {
			logger.Error(err)

			return nil, err
		}

		articles = append(articles, model.Article{
			ID: ID.Int64, Author: author.String, Title: title.String,
			Body: body.String, CreatedAt: createdAt.Time,
		})
	}

	bytes, err := json.Marshal(articles)
	if err != nil {
		logrus.Error(err)

		return nil, err
	}

	if err = ar.cacher.Put(query, bytes); err != nil {
		logger.Error()
	}

	return articles, nil
}
