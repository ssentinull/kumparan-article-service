package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/kumparan-article-service/pkg/model"
	"github.com/ssentinull/kumparan-article-service/pkg/utils"
)

type articleUsecase struct {
	articleRepository model.ArticleRepository
}

func NewArticleUsecase(ar model.ArticleRepository) model.ArticleUsecase {
	return &articleUsecase{
		articleRepository: ar,
	}
}

func (au *articleUsecase) Create(ctx context.Context, article *model.Article) error {
	logger := logrus.WithFields(logrus.Fields{
		"context": utils.Dump(ctx),
		"article": utils.Dump(article),
	})

	err := au.articleRepository.Create(ctx, article)
	if err != nil {
		logger.Error(err)

		return err
	}

	err = au.articleRepository.CalculateVectors(ctx, article)
	if err != nil {
		logger.Error(err)

		return err
	}

	return nil
}

func (au *articleUsecase) Get(ctx context.Context, qp model.QueryBuilder) ([]model.Article, error) {
	articles, err := au.articleRepository.Read(ctx, qp)
	if err != nil {
		logrus.WithField("context", utils.Dump(ctx)).Error(err)

		return nil, err
	}

	return articles, nil
}
