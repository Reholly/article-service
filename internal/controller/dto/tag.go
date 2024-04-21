package dto

import "article-service/internal/domain"

type Tag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func MapToTagEntity(tag Tag) domain.Tag {
	return domain.Tag{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func MapToTagEntities(tags []Tag) []domain.Tag {
	tagEntities := make([]domain.Tag, len(tags))
	for i, v := range tags {
		tagEntities[i] = MapToTagEntity(v)
	}
	return tagEntities
}

func MapToTagDto(tag domain.Tag) Tag {
	return Tag{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func MapToTagDtos(tags []domain.Tag) []Tag {
	tagDtos := make([]Tag, len(tags))
	for i, v := range tags {
		tagDtos[i] = MapToTagDto(v)
	}
	return tagDtos
}
