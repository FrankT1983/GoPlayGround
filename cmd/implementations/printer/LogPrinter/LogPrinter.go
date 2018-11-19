package LogPrinter

import (
	log "github.com/sirupsen/logrus"
)

// PrintToLog implements the IPrinter interface function.
func PrintToLog(toPrint string) {
	log.WithFields(log.Fields{"Origin": "LogPrinter"}).Info(toPrint)
}
