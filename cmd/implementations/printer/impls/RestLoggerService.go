package impls

import (
	"flag"
	"log"
	"reflect"

	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
	restHelper "github.com/FrankT1983/GoPlayGround/restHelper"
	rest "github.com/cseeger-epages/restfool-go"
)

// RestLoggerService is the server side implementatin of a rest rpc bridge
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
	restHelper.AddInterfaceHandlers(api, reflect.TypeOf(new(inf.IPrinter)).Elem(), l.logger)

	// start
	err = api.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
