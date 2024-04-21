package dto

import "errors"

var (
	ErrorEmptyArticleTitle   = errors.New("пустой заголовок статьи")
	ErrorEmptyAuthorUsername = errors.New("пустой заголовок")
	ErrorEmptyTagTitle       = errors.New("пустой заголовок тэга")
)

func ValidateArticleCreation(dto ArticleCreation) error {
	if dto.Title == "" {
		return ErrorEmptyArticleTitle
	}

	return nil
}

func ValidateTagCreation(dto TagCreation) error {
	if dto.Title == "" {
		return ErrorEmptyTagTitle
	}
	return nil
}

func ValidateArticle(dto Article) error {
	if dto.Title == "" {
		return ErrorEmptyArticleTitle
	}

	if dto.AuthorUsername == "" {
		return ErrorEmptyAuthorUsername
	}

	return nil
}
