package domain

import "context"

type ArticleService interface {
	CreateArticle(ctx context.Context, title, text, authorUsername string, tags []Tag) error
	UpdateArticle(ctx context.Context, id int, text, title string, tags []Tag) error
}
