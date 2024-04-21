package app

import (
	"article-service/internal/config"
	"article-service/internal/controller"
	"article-service/internal/repository"
	repoImpl "article-service/internal/repository/implementation"
	"article-service/internal/service"
	serviceImpl "article-service/internal/service/implementation"
	"article-service/lib/adapters/db"
	"context"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	dbAdapter db.PostgresAdapter

	service    service.ServiceManager
	repository repository.RepositoryManager

	config config.Config
}

func NewApplication(config config.Config) *Application {
	return &Application{
		config: config,
	}
}

func (app *Application) Init(ctx context.Context) error {
	app.dbAdapter = db.NewPostgresAdapter()
	_, err := app.dbAdapter.Connect(ctx, app.config.ConnectionString)
	if err != nil {
		return err
	}

	tagRepository := repoImpl.NewTagRepository(app.dbAdapter)
	articleRepository := repoImpl.NewArticleRepository(app.dbAdapter)
	app.repository = repository.NewRepositoryManager(tagRepository, articleRepository)

	articleService := serviceImpl.NewArticleService(app.repository)
	app.service = service.NewServiceManager(articleService)

	migrator := NewMigrator(app.config, app.dbAdapter, app.repository)
	err = migrator.Migrate()
	if err != nil {
		return err
	}
	err = migrator.Seed()
	if err != nil {

	}

	return nil
}

func (app *Application) Run() {
	router := controller.NewRouter(
		app.config.ServerAddress,
		app.config.JwtSecret,
		app.repository,
		app.service,
	)

	router.StartRouting()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
