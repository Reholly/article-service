package response

import (
	"article-service/internal/controller/dto"
	"article-service/internal/domain"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

var (
	SuccessResult = NewAPIResult(http.StatusOK, "OK")
)

type APIResult struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func NewAPIResult(code int, message string) APIResult {
	return APIResult{
		StatusCode: code,
		Message:    message,
	}
}

func CollectError(err error) APIResult {
	log.Println("ERROR:", err)
	switch {
	case errors.Is(err, ErrorInvalidToken):
	case errors.Is(err, ErrorNoToken):
		return NewAPIResult(http.StatusUnauthorized, err.Error())
	case errors.Is(err, ErrorAccessDenied):
		return NewAPIResult(http.StatusForbidden, err.Error())
	case errors.Is(err, domain.ErrorArticleNotFound):
	case errors.Is(err, domain.ErrorTagNotFound):
		return NewAPIResult(http.StatusNotFound, err.Error())
	case errors.Is(err, dto.ErrorEmptyAuthorUsername):
	case errors.Is(err, dto.ErrorEmptyTagTitle):
	case errors.Is(err, ErrorEmptyId):
	case errors.Is(err, ErrorInvalidId):
	case errors.Is(err, dto.ErrorEmptyArticleTitle):
		return NewAPIResult(http.StatusBadRequest, err.Error())
	case errors.Is(err, domain.ErrorEmptyAuthorUsername):
	case errors.Is(err, domain.ErrorEmptyTagTitle):
	case errors.Is(err, domain.ErrorEmptyArticleTitle):
		return NewAPIResult(http.StatusInternalServerError, err.Error())
	}

	return NewAPIResult(http.StatusInternalServerError, "Ошибка сервера, попытайтесь еще раз позже")
}
