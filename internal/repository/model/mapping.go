package model

import "article-service/internal/domain"

func MapArticleModelToEntity(article Article, tags []Tag) domain.Article {
	return domain.Article{
		ID:    article.ID,
		Title: article.Title,
		Text:  article.Text,
		Tags:  MapTagModelsToEntities(tags),

		PublicationDate: article.PublicationDate,
		AuthorUsername:  article.AuthorUsername,
	}
}

func MapToArticleEntities(articles []Article, articleTags []ArticleTagPair, tags []Tag) []domain.Article {
	result := make([]domain.Article, 0, len(articles))

	for _, article := range articles {
		neededTags := make([]Tag, 0)
		for _, articleTag := range articleTags {
			if article.ID == articleTag.ArticleID {
				for _, tag := range tags {
					if tag.ID == articleTag.TagID {
						neededTags = append(neededTags, tag)
						break
					}
				}
			}
		}
		result = append(result, MapArticleModelToEntity(article, neededTags))
	}
	return result
}

func MapTagModelsToEntities(tags []Tag) []domain.Tag {
	res := make([]domain.Tag, 0)
	for _, tag := range tags {
		res = append(res, domain.Tag(tag))
	}
	return res
}
