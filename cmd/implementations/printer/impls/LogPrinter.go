package impls

import (
	log "github.com/sirupsen/logrus"
)

// LogPrinter is an implementation of the IPrinter interface used to print locally.
type LogPrinter struct {
	LoggerName string
}

// PrintToLog implements the IPrinter interface function.
func (l LogPrinter) PrintToLog(toPrint string) {
	log.WithFields(log.Fields{"Origin": "LogPrinter"}).Info(toPrint)
}
