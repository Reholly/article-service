package implementation

import (
	"article-service/internal/domain"
	"article-service/internal/repository"
	"context"
	"time"
)

type ArticleService struct {
	repository repository.RepositoryManager
}

func NewArticleService(repository repository.RepositoryManager) domain.ArticleService {
	return &ArticleService{
		repository: repository,
	}
}

func (s *ArticleService) CreateArticle(ctx context.Context, title, text, authorUsername string, tags []domain.Tag) error {
	tagsFromRepo, err := s.repository.GetAllTags(ctx)
	if err != nil {
		return err
	}

	validTags := make([]domain.Tag, 0)
	for _, tagFromRepo := range tagsFromRepo {
		exist := false
		for _, tag := range tags {
			if tag.ID == tagFromRepo.ID {
				exist = true
				break
			}
		}

		if exist {
			validTags = append(validTags, tagFromRepo)
		}
	}

	newArticle := domain.Article{
		Title:           title,
		Text:            text,
		Tags:            validTags,
		PublicationDate: time.Now(),
		AuthorUsername:  authorUsername,
	}

	err = domain.ValidateArticle(newArticle)
	if err != nil {
		return err
	}

	_, err = s.repository.CreateArticle(ctx, newArticle)
	return err
}

func (s *ArticleService) UpdateArticle(ctx context.Context, id int, text, title string, tags []domain.Tag) error {
	oldArticle, err := s.repository.FindArticleByID(ctx, id)
	if err != nil {
		return err
	}

	oldArticle.Text = text
	oldArticle.Title = title
	oldArticle.Tags = tags

	return s.repository.UpdateArticle(ctx, oldArticle)
}
