package domain

import "context"

type ArticleService interface {
	CreateArticle(ctx context.Context, article Article) error
	FindArticlesByTags(ctx context.Context, tags []Tag) ([]Article, error)
}
