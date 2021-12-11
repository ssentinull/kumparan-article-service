package postgres_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_articleRepo "github.com/ssentinull/kumparan-article-service/pkg/article/repository/postgres"
	"github.com/ssentinull/kumparan-article-service/pkg/model"
	"github.com/ssentinull/kumparan-article-service/pkg/model/mock"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulCreateArticle(t *testing.T) {
	dbStub, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}

	expectedID := int64(1)
	expectedRow := sqlmock.NewRows([]string{"id"}).AddRow(expectedID)
	inputArticle := model.Article{Author: "test author", Title: "test title", Body: "test body"}

	expectedQuery := "INSERT INTO articles (author, title, body, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	expectedPrepare := sqlMock.ExpectPrepare(regexp.QuoteMeta(expectedQuery))
	expectedPrepare.ExpectQuery().WithArgs(inputArticle.Author, inputArticle.Title, inputArticle.Body, mock.AnyTime{}).
		WillReturnRows(expectedRow)

	mockRepo := _articleRepo.NewArticleRepository(dbStub)
	err = mockRepo.Create(context.TODO(), &inputArticle)

	assert.NoError(t, err)
	assert.Equal(t, expectedID, inputArticle.ID)
}

func TestFailedPrepareInCreateArticle(t *testing.T) {
	dbStub, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}

	inputArticle := model.Article{Author: "test author", Title: "test title", Body: "test body"}
	expectedQuery := "INSERT INTO articles (author, title, body, created_at) VALUES ($1, $2, $3, $4)"
	sqlMock.ExpectPrepare(regexp.QuoteMeta(expectedQuery)).WillReturnError(model.ErrInternalServer)
	mockRepo := _articleRepo.NewArticleRepository(dbStub)
	err = mockRepo.Create(context.TODO(), &inputArticle)

	assert.Error(t, err)
	assert.Equal(t, model.ErrInternalServer, err)
}

func TestFailedExecInCreateArticle(t *testing.T) {
	dbStub, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}

	inputArticle := model.Article{Author: "test author", Title: "test title", Body: "test body"}
	expectedQuery := "INSERT INTO articles (author, title, body, created_at) VALUES ($1, $2, $3, $4)"
	expectedPrepare := sqlMock.ExpectPrepare(regexp.QuoteMeta(expectedQuery))
	expectedPrepare.ExpectExec().WithArgs(inputArticle.Author, inputArticle.Title, inputArticle.Body, mock.AnyTime{}).
		WillReturnError(model.ErrBadRequest)
	mockRepo := _articleRepo.NewArticleRepository(dbStub)
	err = mockRepo.Create(context.TODO(), &inputArticle)

	assert.Error(t, err)
	assert.Equal(t, model.ErrBadRequest, err)
}

func TestSuccessfulCalculateVectors(t *testing.T) {
	dbStub, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}

	inputArticle := model.Article{ID: int64(1), Author: "test author", Title: "test title", Body: "test body"}
	expectedQuery := `UPDATE articles SET title_body_vectors = to_tsvector($1 || ' ' || $2),
		author_vectors = to_tsvector($3) WHERE id = $4`
	expectedPrepare := sqlMock.ExpectPrepare(regexp.QuoteMeta(expectedQuery))
	expectedPrepare.ExpectExec().WithArgs(inputArticle.Title, inputArticle.Body, inputArticle.Author, inputArticle.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mockRepo := _articleRepo.NewArticleRepository(dbStub)
	err = mockRepo.CalculateVectors(context.TODO(), &inputArticle)

	assert.NoError(t, err)
}

func TestFailedPrepareInCalculateVectors(t *testing.T) {
	dbStub, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}

	inputArticle := model.Article{ID: int64(1), Author: "test author", Title: "test title", Body: "test body"}
	expectedQuery := `UPDATE articles SET title_body_vectors = to_tsvector($1 || ' ' || $2),
		author_vectors = to_tsvector($3) WHERE id = $4`
	sqlMock.ExpectPrepare(regexp.QuoteMeta(expectedQuery)).WillReturnError(model.ErrInternalServer)

	mockRepo := _articleRepo.NewArticleRepository(dbStub)
	err = mockRepo.CalculateVectors(context.TODO(), &inputArticle)

	assert.Error(t, err)
	assert.Equal(t, model.ErrInternalServer, err)
}

func TestFailedExecInCalculateVectors(t *testing.T) {
	dbStub, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}

	inputArticle := model.Article{ID: int64(1), Author: "test author", Title: "test title", Body: "test body"}
	expectedQuery := `UPDATE articles SET title_body_vectors = to_tsvector($1 || ' ' || $2),
		author_vectors = to_tsvector($3) WHERE id = $4`
	expectedPrepare := sqlMock.ExpectPrepare(regexp.QuoteMeta(expectedQuery))
	expectedPrepare.ExpectExec().WithArgs(inputArticle.Title, inputArticle.Body, inputArticle.Author, inputArticle.ID).
		WillReturnError(model.ErrBadRequest)

	mockRepo := _articleRepo.NewArticleRepository(dbStub)
	err = mockRepo.CalculateVectors(context.TODO(), &inputArticle)

	assert.Error(t, err)
	assert.Equal(t, model.ErrBadRequest, err)
}
