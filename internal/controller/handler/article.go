package handler

import (
	"article-service/internal/controller/dto"
	"article-service/internal/controller/middleware"
	"article-service/internal/controller/response"
	"article-service/internal/repository"
	"article-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	IdParam = "id"
)

type ArticleHandler struct {
	repository repository.RepositoryManager
	service    service.ServiceManager
}

func NewArticleHandler(repository repository.RepositoryManager, service service.ServiceManager) ArticleHandler {
	return ArticleHandler{
		repository: repository,
		service:    service,
	}
}

// CreateArticle godoc
// @Summary Создание статьи
// @Description Этот эндпоинт нужен создания статьи. Принимает Заголовок, Текст (может быть пустым) и тэги Имя автора берется из Claim JWT токена
// @Description Требует админского доступа.
// @Tags article
// @Security Bearer
// @Accept json
// @Produce json
// @Param input body dto.ArticleCreation true "Create"
// @Success 200 {object} response.APIResult
// @Failure 400 {object} response.APIResult
// @Failure 403 {object} response.APIResult
// @Failure 500 {object} response.APIResult
// @Failure 502 {object} response.APIResult
// @Router /api/article [post]
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var newArticle dto.ArticleCreation
	username, _ := c.Get(middleware.UsernameClaim)

	parsedUsername := username.(string)

	if err := c.ShouldBindJSON(&newArticle); err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	if err := dto.ValidateArticleCreation(newArticle); err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	err := h.service.CreateArticle(c.Request.Context(), newArticle.Title, newArticle.Text, parsedUsername, dto.MapToTagEntities(newArticle.Tags))

	if err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	c.JSON(response.SuccessResult.StatusCode, response.SuccessResult)
}

// DeleteArticle godoc
// @Summary Удаление статьи
// @Description Этот эндпоинт нужен для удаления статьи по ID. Требует админского доступа.
// @Tags article
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Article ID"
// @Success 200 {object} response.APIResult
// @Failure 400 {object} response.APIResult
// @Failure 403 {object} response.APIResult
// @Failure 500 {object} response.APIResult
// @Failure 502 {object} response.APIResult
// @Router /api/article/{id} [delete]
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id := c.Param(IdParam)
	if id == "" {
		apiError := response.CollectError(response.ErrorEmptyId)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		apiError := response.CollectError(response.ErrorInvalidId)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}
	err = h.repository.DeleteArticleByID(c.Request.Context(), parsedId)
	if err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	c.JSON(response.SuccessResult.StatusCode, response.SuccessResult)
}

// UpdateArticle godoc
// @Summary Обновление статьи
// @Description Этот эндпоинт нужен для обновления статьи, только текст может быть пустым. Требует админского доступа.
// @Tags article
// @Accept json
// @Security Bearer
// @Produce json
// @Param input body dto.Article true "Update"
// @Success 200 {object} response.APIResult
// @Failure 400 {object} response.APIResult
// @Failure 403 {object} response.APIResult
// @Failure 500 {object} response.APIResult
// @Failure 502 {object} response.APIResult
// @Router /api/article [put]
func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	var updatedArticle dto.Article

	if err := c.ShouldBindJSON(&updatedArticle); err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	err := h.service.UpdateArticle(
		c.Request.Context(),
		updatedArticle.ID,
		updatedArticle.Text,
		updatedArticle.Text,
		dto.MapToTagEntities(updatedArticle.Tags),
	)

	if err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	c.JSON(response.SuccessResult.StatusCode, response.SuccessResult)
}

// GetArticleById godoc
// @Summary Получение статьи по ID
// @Description Этот эндпоинт нужен для получения статьи по ID. Сигнатура описана ниже.
// @Tags article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} dto.Article
// @Failure 400 {object} response.APIResult
// @Failure 403 {object} response.APIResult
// @Failure 500 {object} response.APIResult
// @Failure 502 {object} response.APIResult
// @Router /api/article/{id} [get]
func (h *ArticleHandler) GetArticleById(c *gin.Context) {
	id := c.Param(IdParam)
	if id == "" {
		apiError := response.CollectError(response.ErrorEmptyId)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		apiError := response.CollectError(response.ErrorInvalidId)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	article, err := h.repository.FindArticleByID(c.Request.Context(), parsedId)
	if err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	c.JSON(http.StatusOK, article)
}

// GetAllArticles godoc
// @Summary Получение всех статей
// @Description Этот эндпоинт нужен для получения всех статей.
// @Tags article
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Article
// @Failure 400 {object} response.APIResult
// @Failure 403 {object} response.APIResult
// @Failure 500 {object} response.APIResult
// @Failure 502 {object} response.APIResult
// @Router /api/article [get]
func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	articles, err := h.repository.GetAllArticles(c.Request.Context())
	if err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	c.JSON(http.StatusOK, dto.MapToArticleDtos(articles))
}
