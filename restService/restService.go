package main

import (
	"fmt"

	impl "github.com/FrankT1983/GoPlayGround/cmd/implementations/printer/impls"
)

func main() {
	fmt.Println("Start Rest Logger Service")
	service := impl.NewRestLoggerService("conf/api.conf")
	service.Start()
}
