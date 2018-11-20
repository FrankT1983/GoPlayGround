package factories

import (
	impl "github.com/FrankT1983/GoPlayGround/cmd/implementations/printer/impls"
	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
	restHelper "github.com/FrankT1983/GoPlayGround/restHelper"
)

// BuildLocalPrinter returns an implementation of the IPrinter interface which runs localy
// on the same thread.
func BuildLocalPrinter() inf.IPrinter {
	result := impl.LogPrinter{LoggerName: "LocalPrinter"}
	return result
}

// BuildRestPrinter returns an implementation of the IPrinter interface which runs remotly
// behind a REST interface.
func BuildRestPrinter(info restHelper.ServerInfo) inf.IPrinter {
	result := impl.RestLoggerEndpoint{LoggerName: "RestPrinter", Server: info}
	return result
}
