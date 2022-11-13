package middlewares

import (
	"github.com/Haato3o/poogie/core/utils"
	"github.com/Haato3o/poogie/pkg/log"
	"github.com/gin-gonic/gin"
)

type RequestContext struct {
	UserId string `json:"user_id"`
}

func LogRequest(ctx *gin.Context) {
	log.Info(ctx.Request.URL.Path, &RequestContext{UserId: utils.ExtractUserId(ctx)})
}
