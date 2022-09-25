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

func ElementNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": ELEMENT_NOT_FOUND_MESSAGE})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": INTERNAL_ERROR_MESSAGE})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": UNAUTHORIZED_MESSAGE, "code": common.ErrInvalidCredentials})
}

func Conflict(c *gin.Context, message string, code string) {
	c.JSON(http.StatusConflict, gin.H{"error": message, "code": code})
}

func Ok[T any](c *gin.Context, body T) {
	c.JSON(http.StatusOK, body)
}

func OkZip(c *gin.Context, content []byte) {
	c.Data(http.StatusOK, "application/zip", content)
}
