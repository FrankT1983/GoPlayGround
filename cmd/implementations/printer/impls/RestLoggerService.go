package impls

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	err = api.AddHandler("Print", "POST", "/Print", "Print To Logger", l.printHandler)
	if err != nil {
		log.Fatal(err)
	}

	// start
	err = api.Serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (l RestLoggerService) printHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Someone called print")
	// dont need to cache ?
	w.Header().Set("Cache-Control", "no-store")

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)

	l.logger.PrintToLog(bodyString)

	qs := rest.ParseQueryStrings(r)
	message := fmt.Sprintf("Welcome to restfool take a look at https://%s/help", r.Host)
	msg := rest.Msg{Message: message}
	rest.EncodeAndSend(w, r, qs, msg)
}
