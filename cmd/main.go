package main

import (
	"fmt"

	factory "github.com/FrankT1983/GoPlayGround/cmd/implemenation/printer/PrinterFactory"
)

func main() {
	fmt.Println("Hello World")
	printer = factory.BuildLocalPrinter()
}
