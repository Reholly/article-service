package controller

import (
	"article-service/internal/controller/handler"
	"article-service/internal/controller/middleware"
	"article-service/internal/repository"
	"article-service/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Router struct {
	address    string
	jwtSecret  string
	repository repository.RepositoryManager
	service    service.ServiceManager
}

func NewRouter(address, jwtSecret string, repository repository.RepositoryManager, service service.ServiceManager) *Router {
	return &Router{
		address:    address,
		jwtSecret:  jwtSecret,
		repository: repository,
		service:    service,
	}
}

func (router *Router) StartRouting() {
	articleHandler := handler.NewArticleHandler(router.repository, router.service)
	tagHandler := handler.NewTagHandler(router.repository, router.service)

	g := gin.New()

	g.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := g.Group("/api")

	articlePublic := api.Group("/article")
	{
		articlePublic.GET("/", articleHandler.GetAllArticles)
		articlePublic.GET("/:id", articleHandler.GetArticleById)
	}
	articlePrivate := api.Group("/article", middleware.AuthMiddleware(router.jwtSecret))
	{
		articlePrivate.POST("/", articleHandler.CreateArticle)
		articlePrivate.PUT("/", articleHandler.UpdateArticle)
		articlePrivate.DELETE("/:id", articleHandler.DeleteArticle)
	}

	tagPublic := api.Group("/tag")
	{
		tagPublic.GET("/", tagHandler.GetAllTags)
	}
	tagPrivate := api.Group("/tag", middleware.AuthMiddleware(router.jwtSecret))
	{

		tagPrivate.POST("/", tagHandler.CreateTag)
		tagPrivate.DELETE("/:id", tagHandler.DeleteTag)
	}
	g.GET("/auth/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	go func() { _ = g.Run() }()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
