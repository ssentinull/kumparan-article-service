package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/kumparan-article-service/pkg/model"
)

type ArticleHandlerHTTP struct {
	ArticleUsecase model.ArticleUsecase
}

func NewArticleHandler(e *echo.Echo, au model.ArticleUsecase) {
	handler := &ArticleHandlerHTTP{
		ArticleUsecase: au,
	}

	e.POST("/articles", handler.PostArticle)
}

func (ah *ArticleHandlerHTTP) PostArticle(c echo.Context) error {
	article := new(model.Article)
	err := c.Bind(&article)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = ah.ArticleUsecase.Create(c.Request().Context(), article)
	if err != nil {
		logrus.Error(err)

		return c.JSON(model.GetErrorStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusCreated, article)
}
