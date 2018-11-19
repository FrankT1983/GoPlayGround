package impls

import (
	log "github.com/sirupsen/logrus"
)


type LogPrinter struct {  
    LoggerName   string    
}


// PrintToLog implements the IPrinter interface function.
func (l LogPrinter) PrintToLog(toPrint string) {
	log.WithFields(log.Fields{"Origin": "LogPrinter"}).Info(toPrint)
}
