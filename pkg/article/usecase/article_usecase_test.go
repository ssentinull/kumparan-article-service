package usecase_test

import (
	"context"
	"testing"
	"time"

	_articleUcase "github.com/ssentinull/kumparan-article-service/pkg/article/usecase"
	"github.com/ssentinull/kumparan-article-service/pkg/model"
	_mockArticle "github.com/ssentinull/kumparan-article-service/pkg/model/mock/article"
	_mockQryBuilder "github.com/ssentinull/kumparan-article-service/pkg/model/mock/query_builder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateArticle(t *testing.T) {
	mockRepo := new(_mockArticle.ArticleRepository)
	mockArticle := &model.Article{Author: "test author", Title: "test title", Body: "test body"}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Article")).Return(nil).Once()
		mockRepo.On("CalculateVectors", mock.Anything, mock.AnythingOfType("*model.Article")).Return(nil).Once()
		mockUsecase := _articleUcase.NewArticleUsecase(mockRepo)
		err := mockUsecase.Create(context.TODO(), mockArticle)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-create", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Article")).Return(model.ErrBadRequest).Once()
		mockUsecase := _articleUcase.NewArticleUsecase(mockRepo)
		err := mockUsecase.Create(context.TODO(), mockArticle)

		assert.Error(t, err)
		assert.Equal(t, model.ErrBadRequest, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed-calculate-vectors", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Article")).Return(nil).Once()
		mockRepo.On("CalculateVectors", mock.Anything, mock.AnythingOfType("*model.Article")).
			Return(model.ErrInternalServer).Once()
		mockUsecase := _articleUcase.NewArticleUsecase(mockRepo)
		err := mockUsecase.Create(context.TODO(), mockArticle)

		assert.Error(t, err)
		assert.Equal(t, model.ErrInternalServer, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetArticles(t *testing.T) {
	mockRepo := new(_mockArticle.ArticleRepository)
	mockQueryBuilder := new(_mockQryBuilder.QueryBuilder)
	mockArticles := []model.Article{
		{
			ID: 1, Author: "test author", Title: "test title",
			Body: "test body", CreatedAt: time.Now(),
		},
		{
			ID: 2, Author: "test author", Title: "test title",
			Body: "test body", CreatedAt: time.Now(),
		},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Read", mock.Anything, mock.AnythingOfType("*mocks.QueryBuilder")).
			Return(mockArticles, nil).Once()
		mockUsecase := _articleUcase.NewArticleUsecase(mockRepo)
		articles, err := mockUsecase.Get(context.TODO(), mockQueryBuilder)

		assert.NoError(t, err)
		assert.Len(t, articles, len(mockArticles))
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Read", mock.Anything, mock.AnythingOfType("*mocks.QueryBuilder")).
			Return(nil, model.ErrInternalServer).Once()
		mockUsecase := _articleUcase.NewArticleUsecase(mockRepo)
		articles, err := mockUsecase.Get(context.TODO(), mockQueryBuilder)

		assert.Error(t, err)
		assert.Nil(t, articles)
		assert.Equal(t, model.ErrInternalServer, err)
		mockRepo.AssertExpectations(t)
	})
}
