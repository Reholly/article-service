package domain

import "github.com/pkg/errors"

var (
	ErrorEmptyArticleTitle   = errors.New("пустой заголовок")
	ErrorEmptyAuthorUsername = errors.New("пустое имя автора")
	ErrorEmptyTagTitle       = errors.New("пустой заголовок тэга")
)

func ValidateArticle(article Article) error {
	if article.Title == "" {
		return ErrorEmptyArticleTitle
	}
	if article.AuthorUsername == "" {
		return ErrorEmptyAuthorUsername
	}
	return nil
}

func ValidateTag(tag Tag) error {
	if tag.Title == "" {
		return ErrorEmptyTagTitle
	}

	return nil
}
