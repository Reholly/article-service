package service

import "article-service/internal/domain"

type ServiceManager struct {
	domain.ArticleService
}

func NewServiceManager(article domain.ArticleService) ServiceManager {
	return ServiceManager{
		ArticleService: article,
	}
}
