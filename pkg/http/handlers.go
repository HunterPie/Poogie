package http

import (
	"net/http"

	"github.com/Haato3o/poogie/core/features/common"
	"github.com/gin-gonic/gin"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": NOT_FOUND_MESSAGE})
}

func BadRequest(c *gin.Context, code string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": BAD_REQUEST_MESSAGE, "code": code})
}

func TooLarge(c *gin.Context, code string) {
	c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": FILE_TOO_LARGE, "code": code})
}

func ElementNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": ELEMENT_NOT_FOUND_MESSAGE, "code": common.ErrNotFound})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": INTERNAL_ERROR_MESSAGE, "code": common.ErrInternalError})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": UNAUTHORIZED_MESSAGE, "code": common.ErrInvalidCredentials})
}

func UnauthorizedWithCustomError(c *gin.Context, code string, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message, "code": code})
}

func Conflict(c *gin.Context, message string, code string) {
	c.JSON(http.StatusConflict, gin.H{"error": message, "code": code})
}

func TooManyRequests(c *gin.Context, code string) {
	c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests", "code": code})
}

func Ok[T any](c *gin.Context, body T) {
	c.JSON(http.StatusOK, body)
}

func OkZip(c *gin.Context, content []byte) {
	c.Data(http.StatusOK, "application/zip", content)
}
