package domain

import "time"

type Article struct {
	ID    int
	Title string
	Text  string
	Tags  []Tag

	PublicationDate time.Time
	AuthorUsername  string
}

type Tag struct {
	ID    int
	Title string
}

func NewTag(title string) Tag {
	return Tag{
		Title: title,
	}
}
