package interfaces

// LoggingMessage is a more granualar logging entry
type LoggingMessage struct {
	TypeName      string `json:"type"`
	MessageString string `json:"message"`
}

// IPrinter is an interface to abstact printing to a log.
type IPrinter interface {
	PrintToLog(toPrint string)
	// realy? no function  overloading?
	PrintMessageToLog(toPrint LoggingMessage)
}
