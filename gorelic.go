package gorelic

import (
	"fmt"
	"github.com/LYY/go-workers"
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
	"time"
)

var agent *gorelic.Agent

type GorelicMiddleware struct{}

func (r *GorelicMiddleware) Call(queue string, message *workers.Msg, next func() bool) (acknowledge bool) {
	startTime := time.Now()
	acknowledge = next()
	agent.HTTPTimer.UpdateSince(startTime)
	return
}

func InitNewrelicAgent(license string, appname string, verbose bool) error {

	if license == "" {
		return fmt.Errorf("Please specify NewRelic license")
	}

	agent = gorelic.NewAgent()
	agent.NewrelicLicense = license

	agent.HTTPTimer = metrics.NewTimer()
	agent.CollectHTTPStat = true
	agent.Verbose = verbose

	agent.NewrelicName = appname
	agent.Run()
	return nil
}
