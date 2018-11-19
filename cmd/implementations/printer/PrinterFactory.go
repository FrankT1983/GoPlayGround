package Factories

func BuildLocalPrinter() IPrinter {
	return LogPrinter("foo.log")
}
