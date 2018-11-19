package impls

import (
	log "github.com/sirupsen/logrus"
)

// RestLoggerEndpoint implemnet the IPrinter interface by providing an REST endpoint.
type RestLoggerEndpoint struct {
	LoggerName string
}

// PrintToLog implements the IPrinter interface function.
func (l RestLoggerEndpoint) PrintToLog(toPrint string) {
	log.WithFields(log.Fields{"Origin": "LogPrinter"}).Info(toPrint)
}

func init() {

}
