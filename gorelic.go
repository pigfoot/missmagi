package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
)

var agent *gorelic.Agent

func handler(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	agent.HTTPTimer.UpdateSince(startTime)
}

func InitNewrelicAgent(license string, appname string, verbose bool) (func(*gin.Context), error) {

	if license == "" {
		return nil, fmt.Errorf("Please specify NewRelic license")
	}

	agent = gorelic.NewAgent()
	agent.NewrelicLicense = license

	agent.HTTPTimer = metrics.NewTimer()
	agent.CollectHTTPStat = true
	agent.Verbose = verbose

	agent.NewrelicName = appname
	agent.Run()

	return handler, nil
}
