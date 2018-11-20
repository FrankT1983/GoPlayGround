package impls

import (
	"log"

	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
	restHelper "github.com/FrankT1983/GoPlayGround/restHelper"
)

// RestLoggerEndpoint implemnet the IPrinter interface by providing an REST endpoint.
type RestLoggerEndpoint struct {
	LoggerName string
	Server     restHelper.ServerInfo
}

// PrintToLog implements the IPrinter interface function.
func (l RestLoggerEndpoint) PrintToLog(toPrint string) {
	// send rest request
	restHelper.CallStringFunctionOferRest(l.Server, "PrintToLog", toPrint)
}

// PrintMessageToLog implements the IPrinter interface function.
func (l RestLoggerEndpoint) PrintMessageToLog(toPrint inf.LoggingMessage) {
	log.WithFields(log.Fields{"Type": toPrint.TypeName}).Info(toPrint.MessageString)
}
