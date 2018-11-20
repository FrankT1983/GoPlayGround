package interfaces

// LoggingMessage is a more granualar logging entry
type LoggingMessage struct {
	TypeName      string
	MessageString string
}

// IPrinter is an interface to abstact printing to a log.
type IPrinter interface {
	PrintToLog(toPrint string)
	// realy? no function  overloading?
	PrintMessageToLog(toPrint LoggingMessage)
}
