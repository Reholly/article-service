package model

type ArticleTagPair struct {
	ArticleID int `db:"article_id"`
	TagID     int `db:"tag_id"`
}
