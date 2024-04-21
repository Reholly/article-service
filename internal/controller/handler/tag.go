package handler

import (
	"article-service/internal/controller/dto"
	"article-service/internal/controller/response"
	"article-service/internal/domain"
	"article-service/internal/repository"
	"article-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagHandler struct {
	service    service.ServiceManager
	repository repository.RepositoryManager
}

func NewTagHandler(repository repository.RepositoryManager, service service.ServiceManager) TagHandler {
	return TagHandler{
		service:    service,
		repository: repository,
	}
}

// CreateTag godoc
// @Summary Создание тэга
// @Description Этот эндпоинт нужен для получения создания тэга, принимает только название. Если тэг уже есть, то вернет ошибку.
// @Tags tag
// @Accept json
// @Security Bearer
// @Param input body dto.TagCreation true "Create"
// @Produce json
// @Success 200 {object} response.APIResult
// @Failure 400 {object} response.APIResult
// @Failure 403 {object} response.APIResult
// @Failure 500 {object} response.APIResult
// @Failure 502 {object} response.APIResult
// @Router /api/tag [post]
func (h *TagHandler) CreateTag(c *gin.Context) {
	var tag dto.TagCreation

	if err := c.ShouldBindJSON(&tag); err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError.Message)
		return
	}

	if err := dto.ValidateTagCreation(tag); err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	err := h.repository.CreateTag(c.Request.Context(), domain.NewTag(tag.Title))
	if err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}
	c.JSON(response.SuccessResult.StatusCode, response.SuccessResult)
}

// DeleteTag godoc
// @Summary Удаления тэга
// @Description Этот эндпоинт нужен для получения удаления тэга, принимает только ID.
// @Tags tag
// @Accept json
// @Security Bearer
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200 {object} response.APIResult
// @Failure 400 {object} response.APIResult
// @Failure 403 {object} response.APIResult
// @Failure 500 {object} response.APIResult
// @Failure 502 {object} response.APIResult
// @Router /api/tag/{id} [delete]
func (h *TagHandler) DeleteTag(c *gin.Context) {
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

	err = h.repository.DeleteTagByID(c.Request.Context(), parsedId)
	if err != nil {
		apiError := response.CollectError(response.ErrorInvalidId)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	c.JSON(response.SuccessResult.StatusCode, response.SuccessResult)
}

// GetAllTags godoc
// @Summary Получение всех тэгов.
// @Description Этот эндпоинт нужен для получения всех тэгов.
// @Tags tag
// @Accept json
// @Produce json
// @Success 200 {object} []dto.Tag
// @Failure 400 {object} response.APIResult
// @Failure 403 {object} response.APIResult
// @Failure 500 {object} response.APIResult
// @Failure 502 {object} response.APIResult
// @Router /api/tag [get]
func (h *TagHandler) GetAllTags(c *gin.Context) {
	tags, err := h.repository.GetAllTags(c.Request.Context())
	if err != nil {
		apiError := response.CollectError(err)
		c.AbortWithStatusJSON(apiError.StatusCode, apiError)
		return
	}

	c.JSON(http.StatusOK, dto.MapToTagDtos(tags))
}
