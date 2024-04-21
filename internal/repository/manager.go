package repository

import "article-service/internal/domain"

type RepositoryManager struct {
	domain.TagRepository
	domain.ArticleRepository
}

func NewRepositoryManager(tag domain.TagRepository, article domain.ArticleRepository) RepositoryManager {
	return RepositoryManager{
		TagRepository:     tag,
		ArticleRepository: article,
	}
}
