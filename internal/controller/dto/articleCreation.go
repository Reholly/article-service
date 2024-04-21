package dto

type ArticleCreation struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Tags  []Tag  `json:"tags"`
}
