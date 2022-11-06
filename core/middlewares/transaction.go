package middlewares

import (
	"github.com/Haato3o/poogie/core/tracing"
	"github.com/Haato3o/poogie/core/utils"
	"github.com/gin-gonic/gin"
)

const (
	ClientId        = "X-Client-Id"
	AppVersion      = "X-App-Version"
	HunterPieClient = "X-HunterPie-Client"
	ConnectingIp    = "Do-Connecting-Ip"
)

func TransactionMiddleware(ctx *gin.Context) {
	txn := tracing.FromContext(ctx)

	if txn == nil {
		return
	}

	txn.AddProperty("client-id", ctx.Request.Header.Get(ClientId))
	txn.AddProperty("app-version", ctx.Request.Header.Get(AppVersion))
	txn.AddProperty("client-type", ctx.Request.Header.Get(HunterPieClient))
	txn.AddProperty("supporter", utils.ExtractIsSupporter(ctx))
	txn.AddProperty("user-agent", ctx.Request.UserAgent())
	txn.AddProperty("client-ip", ctx.Request.Header.Get(ConnectingIp))
	txn.AddProperty("user_id", utils.ExtractUserId(ctx))
}
