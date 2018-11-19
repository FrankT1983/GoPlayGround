package interfaces

// IPrinter is an interface to abstact printing to a log.
type IPrinter interface {
	PrintToLog(toPrint string)
}
