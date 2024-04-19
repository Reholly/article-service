package domain

import (
	"context"
	"errors"
)

var (
	ErrorArticleNotFound = errors.New("статья не найдена")
	ErrorTagNotFound     = errors.New("тэг не найден")
)

type ArticleRepository interface {
	FindArticleByID(ctx context.Context, id int) (Article, error)
	GetAllArticles(ctx context.Context) ([]Article, error)

	CreateArticle(ctx context.Context, article Article) error
	DeleteArticleByID(ctx context.Context, id int) error
	UpdateArticle(ctx context.Context, article Article) error

	AddTagToArticle(ctx context.Context, articleId, tagId int) error
	RemoveTagFromArticle(ctx context.Context, articleId, tagId int) error
}

type TagRepository interface {
	GetAllTags(ctx context.Context) ([]Tag, error)
	CreateTag(ctx context.Context, tag Tag) error
	DeleteTagByID(ctx context.Context, id int) error
}
