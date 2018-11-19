package main

import (
	"fmt"

	printerFactory "github.com/FrankT1983/GoPlayGround/cmd/implementations/printer"
)

func main() {
	fmt.Println("Hello World")
	printer := printerFactory.BuildLocalPrinter()
	printer.PrintToLog("Foo")

	fmt.Println("Print 2")
	printer2 := printerFactory.BuildRestPrinter()
	printer2.PrintToLog("Foo over rest")

}
