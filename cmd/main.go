package main

import (
	"fmt"
	"github.com/FrankT1983/GoPlayGround/cmd/implementations/printer"
)

func main() {
	fmt.Println("Hello World")
	PrinterFacotry factory;
	printer := factory.BuildLocalPrinter()
	printer.PrintToLog("Foo")
}


