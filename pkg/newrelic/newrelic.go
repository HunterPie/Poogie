package newrelic

import (
	"time"

	"github.com/Haato3o/poogie/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type NewRelicTracingEngine struct {
	App     *newrelic.Application
	appName string
}

func New(env, licenseKey string, isEnabled bool) *NewRelicTracingEngine {
	nr := &NewRelicTracingEngine{
		appName: "poogie-api:" + env,
	}

	var err error
	nr.App, err = newrelic.NewApplication(
		newrelic.ConfigAppName(nr.appName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigEnabled(isEnabled),
		newrelic.ConfigDistributedTracerEnabled(true),
		nrzap.ConfigLogger(log.ZapLogger.Named("newrelic")),
	)

	if err != nil {
		log.Error("failed to create New Relic agent instance", err)
		return nr
	}

	if err = nr.App.WaitForConnection(5 * time.Second); err != nil {
		log.Error("failed to connect to NewRelic", err)
		return nr
	}

	log.Info("NewRelic initialized")

	return nr
}

func (e *NewRelicTracingEngine) SetupTracingMiddleware(router *gin.Engine) {
	router.Use(nrgin.Middleware(e.App))
}
