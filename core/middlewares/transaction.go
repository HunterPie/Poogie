package middlewares

import (
	"github.com/Haato3o/poogie/core/tracing"
	"github.com/gin-gonic/gin"
)

const (
	CLIENT_ID        = "X-Client-Id"
	APP_VERSION      = "X-App-Version"
	SUPPORTER        = "X-Supporter-Token"
	HUNTERPIE_CLIENT = "X-HunterPie-Client"
	CONNECTING_IP    = "Do-Connecting-Ip"
)

func TransactionMiddleware(ctx *gin.Context) {
	txn := tracing.FromContext(ctx)

	if txn == nil {
		return
	}

	txn.AddProperty("client-id", ctx.Request.Header.Get(CLIENT_ID))
	txn.AddProperty("app-version", ctx.Request.Header.Get(APP_VERSION))
	txn.AddProperty("client-type", ctx.Request.Header.Get(HUNTERPIE_CLIENT))
	txn.AddProperty("supporter", ctx.Request.Header.Get(SUPPORTER) != "")
	txn.AddProperty("user-agent", ctx.Request.UserAgent())
	txn.AddProperty("client-ip", ctx.Request.Header.Get(CONNECTING_IP))
}
