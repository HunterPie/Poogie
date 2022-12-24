package utils

import (
	"github.com/Haato3o/poogie/core/tracing"
	"github.com/Haato3o/poogie/pkg/log"
	"github.com/gin-gonic/gin"
	"strconv"
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

func ExtractIsSupporter(ctx *gin.Context) bool {
	ok, err := strconv.ParseBool(ctx.GetHeader("X-Transformed-Is-Supporter"))

	if err != nil {
		return false
	}

	return ok
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

func DeserializeBody[T any](ctx *gin.Context, body *T, validators ...func(*T) (bool, bool)) (bool, bool) {
	err := ctx.BindJSON(body)

	txn := tracing.FromContext(ctx)

	if err != nil {
		log.Error("failed to deserialize body", err)
		txn.AddProperty("error_message", err)
		return false, false
	}

	for _, validator := range validators {
		success, handled := validator(body)
		if !success {
			return false, handled
		}
	}

	return true, false
}
