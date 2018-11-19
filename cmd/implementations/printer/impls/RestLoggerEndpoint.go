package impls

import (
	restHelper "github.com/FrankT1983/GoPlayGround/restHelper"
)

// RestLoggerEndpoint implemnet the IPrinter interface by providing an REST endpoint.
type RestLoggerEndpoint struct {
	LoggerName string
}

// PrintToLog implements the IPrinter interface function.
func (l RestLoggerEndpoint) PrintToLog(toPrint string) {
	// send rest request
	restHelper.CallStringFunctionOferRest("PrintToLog", toPrint)
}
