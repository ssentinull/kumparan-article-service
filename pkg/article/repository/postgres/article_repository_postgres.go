package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/kumparan-article-service/pkg/model"
	"github.com/ssentinull/kumparan-article-service/pkg/utils"
)

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) model.ArticleRepository {
	return &articleRepository{
		db: db,
	}
}

func (ar *articleRepository) Create(ctx context.Context, article *model.Article) error {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"article": utils.Dump(article),
	})

	query := "INSERT INTO articles (author, title, body, created_at) VALUES ($1, $2, $3, $4)"
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

	now := time.Now().In(jkt)
	_, err = stmt.ExecContext(ctx, article.Author, article.Title, article.Body, now)
	if err != nil {
		logger.Error(err)
		err = model.ErrBadRequest

		return err
	}

	article.CreatedAt = now

	return nil
}
