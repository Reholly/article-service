package dto

import (
	"article-service/internal/domain"
	"time"
)

type Article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Tags  []Tag  `json:"tags"`

	PublicationDate time.Time `json:"publication_date"`
	AuthorUsername  string    `json:"author"`
}

func MapToArticleDto(article domain.Article) Article {
	return Article{
		ID:              article.ID,
		Title:           article.Title,
		Text:            article.Text,
		Tags:            MapToTagDtos(article.Tags),
		PublicationDate: article.PublicationDate,
		AuthorUsername:  article.AuthorUsername,
	}
}

func MapToArticleDtos(articles []domain.Article) []Article {
	articleEntities := make([]Article, len(articles))
	for i, v := range articles {
		articleEntities[i] = MapToArticleDto(v)
	}
	return articleEntities
}
