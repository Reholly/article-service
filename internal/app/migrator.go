package app

import (
	"article-service/internal/config"
	"article-service/internal/domain"
	"article-service/internal/repository"
	"article-service/lib/adapters/db"
	"context"
	"errors"
	"github.com/pressly/goose"
	"log"
)

var (
	ErrorConnectionIsNil = errors.New("[ Migrator ] adapter connection is nil")
)

type Migrator struct {
	config     config.Config
	adapter    db.PostgresAdapter
	repository repository.RepositoryManager
}

func NewMigrator(
	config config.Config,
	adapter db.PostgresAdapter,
	repository repository.RepositoryManager,
) Migrator {
	return Migrator{
		config:     config,
		adapter:    adapter,
		repository: repository,
	}
}

func (migrator *Migrator) Migrate() error {
	connection := migrator.adapter.GetConnection()
	if connection == nil {
		return ErrorConnectionIsNil
	}
	if err := goose.SetDialect(migrator.config.DriverName); err != nil {
		return err
	}

	if err := goose.Up(connection.DB, migrator.config.MigrationsDir); err != nil {
		return err
	}
	return nil
}

func (migrator *Migrator) Seed() error {
	tags := []domain.Tag{
		domain.NewTag("binary tree"),
		domain.NewTag("DFS"),
		domain.NewTag("BFS"),
		domain.NewTag("sliding window"),
		domain.NewTag("matrix "),
		domain.NewTag("sort"),
		domain.NewTag("hash table"),
		domain.NewTag("DP"),
		domain.NewTag("two pointers"),
		domain.NewTag("graph"),
		domain.NewTag("stack"),
		domain.NewTag("queue"),
	}

	ctx := context.Background()
	for _, v := range tags {
		err := migrator.repository.CreateTag(ctx, v)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
