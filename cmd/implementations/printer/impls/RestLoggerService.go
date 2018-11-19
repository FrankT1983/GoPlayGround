package impls

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
	rest "github.com/cseeger-epages/restfool-go"
)

type RestLoggerService struct {
	logger     inf.IPrinter
	configPaht string
}

// NewRestLoggerService creates a new service
func NewRestLoggerService(myConfigPaht string) RestLoggerService {
	result := RestLoggerService{
		logger:     LogPrinter{LoggerName: "foo"},
		configPaht: myConfigPaht,
	}
	return result
}

// Start starts the service
func (l RestLoggerService) Start() {
	confFile := flag.String("c", l.configPaht, "path to config ile")

	flag.Parse()
	// realy would like to use the Factory
	l.logger = LogPrinter{LoggerName: "foo"}

	// initialize rest api using conf file
	api, err := rest.New(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	// add handler
	err = api.AddHandler("Print", "GET", "/", "Print To Logger", l.index)
	if err != nil {
		log.Fatal(err)
	}

	// start
	err = api.Serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (l RestLoggerService) index(w http.ResponseWriter, r *http.Request) {
	// dont need to cache ?
	w.Header().Set("Cache-Control", "no-store")

	l.logger.PrintToLog("Foo")

	qs := rest.ParseQueryStrings(r)
	message := fmt.Sprintf("Welcome to restfool take a look at https://%s/help", r.Host)
	msg := rest.Msg{Message: message}
	rest.EncodeAndSend(w, r, qs, msg)
}
