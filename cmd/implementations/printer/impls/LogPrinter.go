package impls

import (
	log "github.com/sirupsen/logrus"

	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
)

// LogPrinter is an implementation of the IPrinter interface used to print locally.
type LogPrinter struct {
	LoggerName string
}

// PrintToLog implements the IPrinter interface function.
func (l LogPrinter) PrintToLog(toPrint string) {
	log.WithFields(log.Fields{"Origin": l.LoggerName}).Info(toPrint)
}

// PrintMessageToLog implements the IPrinter interface function.
func (l LogPrinter) PrintMessageToLog(toPrint inf.LoggingMessage) {
	log.WithFields(log.Fields{"Origin": l.PrintMessageToLog, "Type": toPrint.TypeName}).Info(toPrint.MessageString)
}
