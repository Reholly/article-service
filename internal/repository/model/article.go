package model

import (
	"time"
)

type Article struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Text  string `db:"text"`

	PublicationDate time.Time `db:"publication_date"`
	AuthorUsername  string    `db:"author_username"`
}
