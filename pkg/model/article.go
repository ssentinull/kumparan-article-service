package model

import (
	"context"
	"time"
)

type Article struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleUsecase interface {
	Create(context.Context, *Article) error
	Get(context.Context, QueryParam) ([]Article, error)
}

type ArticleRepository interface {
	CalculateVectors(context.Context, *Article) error
	Create(context.Context, *Article) error
	Read(context.Context, QueryParam) ([]Article, error)
}
