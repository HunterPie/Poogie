package utils

import (
	"github.com/gin-gonic/gin"
)

func ExtractSupporterToken(ctx *gin.Context) string {
	return ctx.GetHeader("X-Supporter-Token")
}

func ExtractClientId(ctx *gin.Context) string {
	return ctx.GetHeader("X-Client-Id")
}

func ExtractUserId(ctx *gin.Context) string {
	return ctx.GetHeader("X-Transformed-User-Id")
}

func DeserializeHeaders[T any](ctx *gin.Context, header *T, validators ...func(*T) bool) bool {
	err := ctx.BindHeader(header)

	for _, validator := range validators {
		if !validator(header) {
			return false
		}

	}

	return err == nil
}

func DeserializeBody[T any](ctx *gin.Context, body *T, validators ...func(*T) bool) bool {
	err := ctx.BindJSON(body)

	for _, validator := range validators {
		if !validator(body) {
			return false
		}
	}

	return err == nil
}
