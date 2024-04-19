package model

type Tag struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
