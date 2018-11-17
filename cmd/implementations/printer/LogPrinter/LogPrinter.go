package LogPrinter

import (
  log "github.com/sirupsen/logrus"
)

PrintToLog(toPrint string){
    log.WithFields(log.Fields{"Origin": "LogPrinter",}).Info(toPrint)
}
