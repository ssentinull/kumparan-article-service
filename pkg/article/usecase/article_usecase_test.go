package usecase_test

import (
	"context"
	"testing"

	_articleUcase "github.com/ssentinull/kumparan-article-service/pkg/article/usecase"
	"github.com/ssentinull/kumparan-article-service/pkg/model"
	_mock "github.com/ssentinull/kumparan-article-service/pkg/model/mock/article"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateArticle(t *testing.T) {
	mockRepo := new(_mock.ArticleRepository)
	mockArticle := &model.Article{Author: "test author", Title: "test title", Body: "test body"}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Article")).Return(nil).Once()
		mockUsecase := _articleUcase.NewArticleUsecase(mockRepo)
		err := mockUsecase.Create(context.TODO(), mockArticle)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Article")).Return(model.ErrBadRequest).Once()
		mockUsecase := _articleUcase.NewArticleUsecase(mockRepo)
		err := mockUsecase.Create(context.TODO(), mockArticle)

		assert.Error(t, err)
		assert.Equal(t, model.ErrBadRequest, err)
		mockRepo.AssertExpectations(t)
	})
}
