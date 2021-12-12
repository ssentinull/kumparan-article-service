package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/labstack/echo/v4"
	_articleHndlr "github.com/ssentinull/kumparan-article-service/pkg/article/handler/http"
	"github.com/ssentinull/kumparan-article-service/pkg/model"
	_mock "github.com/ssentinull/kumparan-article-service/pkg/model/mock/article"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulPostArticle(t *testing.T) {
	var mockArticle model.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	j, err := json.Marshal(mockArticle)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(echo.POST, "/articles", strings.NewReader(string(j)))
	assert.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e := echo.New()
	c := e.NewContext(req, rec)

	mockUsecase := new(_mock.ArticleUsecase)
	mockUsecase.On("Create", mock.Anything, mock.AnythingOfType("*model.Article")).Return(nil)
	handler := _articleHndlr.ArticleHandlerHTTP{ArticleUsecase: mockUsecase}
	err = handler.PostArticle(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestFailedBindInPostArticle(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(echo.POST, "/articles", strings.NewReader(string("faulty-request-body")))
	assert.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e := echo.New()
	c := e.NewContext(req, rec)

	mockUsecase := new(_mock.ArticleUsecase)
	handler := _articleHndlr.ArticleHandlerHTTP{ArticleUsecase: mockUsecase}
	err = handler.PostArticle(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestFailedUsecaseInPostArticle(t *testing.T) {
	var mockArticle model.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	j, err := json.Marshal(mockArticle)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(echo.POST, "/articles", strings.NewReader(string(j)))
	assert.NoError(t, err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e := echo.New()
	c := e.NewContext(req, rec)

	mockUsecase := new(_mock.ArticleUsecase)
	mockUsecase.On("Create", mock.Anything, mock.AnythingOfType("*model.Article")).Return(model.ErrInternalServer)
	handler := _articleHndlr.ArticleHandlerHTTP{ArticleUsecase: mockUsecase}
	err = handler.PostArticle(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestSuccessfulFetchArticles(t *testing.T) {
	params := make(url.Values)
	params.Set("query", "science")
	params.Set("author", "john")

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(echo.GET, "/articles?"+params.Encode(), strings.NewReader(""))
	assert.NoError(t, err)

	e := echo.New()
	c := e.NewContext(req, rec)

	var article model.Article
	err = faker.FakeData(&article)
	assert.NoError(t, err)

	articles := []model.Article{article}
	mockUsecase := new(_mock.ArticleUsecase)
	mockUsecase.On("Get", mock.Anything, mock.AnythingOfType("query_param.QueryParam")).Return(articles, nil)
	handler := _articleHndlr.ArticleHandlerHTTP{ArticleUsecase: mockUsecase}
	err = handler.FetchArticles(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestFailedUsecaseInFetchArticles(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(echo.GET, "/articles", strings.NewReader(""))
	assert.NoError(t, err)

	e := echo.New()
	c := e.NewContext(req, rec)

	mockUsecase := new(_mock.ArticleUsecase)
	mockUsecase.On("Get", mock.Anything, mock.AnythingOfType("query_param.QueryParam")).
		Return(nil, model.ErrInternalServer)
	handler := _articleHndlr.ArticleHandlerHTTP{ArticleUsecase: mockUsecase}
	err = handler.FetchArticles(c)

	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.AssertExpectations(t)
}
