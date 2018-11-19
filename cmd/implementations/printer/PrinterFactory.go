package factories

import (
	impl "github.com/FrankT1983/GoPlayGround/cmd/implementations/printer/impls"
	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
)

// BuildLocalPrinter returns an implementation of the IPrinter interface which runs localy
// on the same thread.
func BuildLocalPrinter() inf.IPrinter {
	result := impl.LogPrinter{LoggerName: "foo"}
	return result
}
