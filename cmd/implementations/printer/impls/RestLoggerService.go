package impls

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
	rest "github.com/cseeger-epages/restfool-go"
)

var logger inf.IPrinter

func main() {
	confFile := flag.String("c", "conf/api.conf", "path to config ile")

	flag.Parse()
	// realy would like to use the Factory
	logger = LogPrinter{LoggerName: "foo"}

	// initialize rest api using conf file
	api, err := rest.New(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	// add handler
	err = api.AddHandler("Print", "GET", "/", "Print To Logger", index)
	if err != nil {
		log.Fatal(err)
	}

	// start
	err = api.Serve()
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// dont need to cache ?
	w.Header().Set("Cache-Control", "no-store")

	logger.PrintToLog("Foo")

	qs := rest.ParseQueryStrings(r)
	message := fmt.Sprintf("Welcome to restfool take a look at https://%s/help", r.Host)
	msg := rest.Msg{Message: message}
	rest.EncodeAndSend(w, r, qs, msg)
}
