package implementation

import (
	"article-service/internal/domain"
	"article-service/internal/repository/model"
	"article-service/lib/adapters/db"
	"context"
	sql2 "database/sql"
	"errors"
)

type TagRepository struct {
	db db.PostgresAdapter
}

func NewTagRepository(adapter db.PostgresAdapter) domain.TagRepository {
	return &TagRepository{db: adapter}
}

func (r *TagRepository) GetAllTags(ctx context.Context) ([]domain.Tag, error) {
	query := `SELECT id, title FROM tag`

	var tags []model.Tag
	err := r.db.Query(ctx, &tags, query)
	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return nil, domain.ErrorTagNotFound
		}
		return nil, err
	}
	return model.MapTagModelsToEntities(tags), nil
}

func (r *TagRepository) CreateTag(ctx context.Context, tag domain.Tag) error {
	sql := `INSERT INTO tag(id, title) ValUES($1,$2)`
	return r.db.Execute(ctx, sql, tag.ID, tag.Title)
}

func (r *TagRepository) DeleteTagByID(ctx context.Context, id int) error {
	sql := `DELETE FROM tag WHERE id = $1`
	return r.db.Execute(ctx, sql, id)
}
