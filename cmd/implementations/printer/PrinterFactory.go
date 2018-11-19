package factories

import (
	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
	impl "github.com/FrankT1983/GoPlayGround/cmd/implementations/printer/impls"
)


// BuildLocalPrinter returns an implementation of the IPrinter interface which runs localy
// on the same thread.
func BuildLocalPrinter() inf.IPrinter {
	result := impl.LogPrinter{"foo"}
	return result;
}
