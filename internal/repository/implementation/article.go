package implementation

import (
	"article-service/internal/domain"
	"article-service/internal/repository/model"
	"article-service/lib/adapters/db"
	"context"
	sql2 "database/sql"
	"errors"
)

type ArticleRepository struct {
	db db.PostgresAdapter
}

func NewArticleRepository(adapter db.PostgresAdapter) domain.ArticleRepository {
	return &ArticleRepository{db: adapter}
}

func (r *ArticleRepository) FindArticleByID(ctx context.Context, id int) (domain.Article, error) {
	sql := `SELECT id, title, text, publication_date, author_username
				FROM article
				WHERE id = $1`

	var article model.Article
	err := r.db.QueryRow(ctx, &article, sql, id)
	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return domain.Article{}, domain.ErrorArticleNotFound
		}
		return domain.Article{}, err
	}

	sql = `SELECT id, title
				FROM tag
				WHERE id IN (
					SELECT tag_id 	
						FROM article_tag 
						WHERE article_id = $1)`

	var articleTagPairs []model.Tag
	err = r.db.Query(ctx, &articleTagPairs, sql, id)
	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return model.MapArticleModelToEntity(article, articleTagPairs), nil
		}
		return domain.Article{}, err
	}

	return model.MapArticleModelToEntity(article, articleTagPairs), nil
}

func (r *ArticleRepository) GetAllArticles(ctx context.Context) ([]domain.Article, error) {
	articleQuery := `SELECT id, title, text, publication_date, author_username
						FROM article`
	var articles []model.Article
	err := r.db.Query(ctx, &articles, articleQuery)
	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, domain.ErrorArticleNotFound
		}
		return nil, err
	}

	tagsQuery := `SELECT id, title
						FROM tag`
	var tags []model.Tag
	err = r.db.Query(ctx, &tags, tagsQuery)
	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, domain.ErrorTagNotFound
		}
		return nil, err
	}

	articleTagQuery := `SELECT tag_id, article_id
							FROM article_tag`
	var articleTags []model.ArticleTagPair
	err = r.db.Query(ctx, &articleTags, articleTagQuery)
	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, domain.ErrorTagNotFound
		}
		return nil, err
	}

	result := model.MapToArticleEntities(articles, articleTags, tags)
	return result, err
}

func (r *ArticleRepository) CreateArticle(ctx context.Context, article domain.Article) (int, error) {
	sql := `INSERT INTO article(text, title, author_username, publication_date)  VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.db.ExecuteAndGet(ctx, &id, sql, article.Text, article.Title, article.AuthorUsername, article.PublicationDate)
	if err != nil {
		return 0, err
	}

	for _, tag := range article.Tags {
		err = r.AddTagToArticle(ctx, id, tag.ID)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *ArticleRepository) DeleteArticleByID(ctx context.Context, id int) error {
	sql := `DELETE FROM article WHERE id = $1`
	err := r.db.Execute(ctx, sql, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ArticleRepository) UpdateArticle(ctx context.Context, article domain.Article) error {
	sql := `UPDATE article 
				SET text=$1, title=$2, author_username=$3, publication_date=$4 
				WHERE id = $5`
	err := r.db.Execute(ctx, sql, article.Text, article.Title, article.AuthorUsername, article.PublicationDate, article.ID)
	if err != nil {
		return err
	}

	sql = `DELETE FROM article_tag WHERE article_id = $1`
	err = r.db.Execute(ctx, sql, article.ID)
	if err != nil {
		return err
	}

	for _, tag := range article.Tags {
		err = r.AddTagToArticle(ctx, article.ID, tag.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ArticleRepository) AddTagToArticle(ctx context.Context, articleId, tagId int) error {
	sql := `INSERT INTO 
    			article_tag(article_id, tag_id) 
				VALUES ($1, $2)`
	err := r.db.Execute(ctx, sql, articleId, tagId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ArticleRepository) RemoveTagFromArticle(ctx context.Context, articleId, tagId int) error {
	sql := `DELETE FROM article_tag
				WHERE article_id = $1 AND tag_id = $2`
	err := r.db.Execute(ctx, sql, articleId, tagId)
	if err != nil {
		return err
	}
	return nil
}
