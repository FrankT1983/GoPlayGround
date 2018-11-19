package factories

import (
	"github.com/FrankT1983/GoPlayGround/cmd/interface"
	"github.com/FrankT1983/GoPlayGround/cmd/implementations/printer/LogPrinter"
)

// BuildLocalPrinter returns an implementation of the IPrinter interface which runs localy
// on the same thread.
func BuildLocalPrinter() IPrinterIPrinter {
	return factory.LogPrinter("foo.log")
}
