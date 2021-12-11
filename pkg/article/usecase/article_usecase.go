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
	err := au.articleRepository.Create(ctx, article)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"context": utils.Dump(ctx),
			"article": utils.Dump(article),
		}).Error(err)

		return err
	}

	return nil
}

func (au *articleUsecase) Get(ctx context.Context) ([]model.Article, error) {
	articles, err := au.articleRepository.Read(ctx)
	if err != nil {
		logrus.WithField("context", utils.Dump(ctx)).Error(err)

		return nil, err
	}

	return articles, nil
}
