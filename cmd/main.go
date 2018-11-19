package main

import (
	"fmt"

	printerFactory "github.com/FrankT1983/GoPlayGround/cmd/implementations/printer"
)

func main() {
	fmt.Println("Hello World")
	printer := printerFactory.BuildLocalPrinter()
	printer.PrintToLog("Foo")
}
