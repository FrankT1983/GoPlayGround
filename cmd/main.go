package main

import (
	"fmt"

	printerFactory "github.com/FrankT1983/GoPlayGround/cmd/implementations/printer"
	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
	restHelper "github.com/FrankT1983/GoPlayGround/restHelper"
)

func main() {
	fmt.Println("Hello World")
	printer := printerFactory.BuildLocalPrinter()
	printer.PrintToLog("Foo")

	serverInfo := restHelper.ServerInfo{
		Server:   "127.0.0.1",
		Port:     9443,
		User:     "testuser",
		Password: "testpass"}
	fmt.Println("Print 2")
	printer2 := printerFactory.BuildRestPrinter(serverInfo)
	printer2.PrintToLog("Foo over rest")

	message = inf.LoggingMessage{MessageString: "More Complex", TypeName: "TestMessage"}
	printer2.PrintToLog("Foo over rest")
}
