package impls

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"reflect"

	inf "github.com/FrankT1983/GoPlayGround/cmd/interface"
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
	l.AddInterfaceHandlers(api, reflect.TypeOf(new(inf.IPrinter)).Elem(), l.logger)

	// start
	err = api.Serve()
	if err != nil {
		log.Fatal(err)
	}
}

// AddInterfaceHandlers generate the handlers for a interface ... code generatio would be the way to go here
func (l RestLoggerService) AddInterfaceHandlers(api rest.RestAPI, interfaceType reflect.Type, impl interface{}) {

	for i := 0; i < interfaceType.NumMethod(); i++ {
		methodName := interfaceType.Method(i).Name

		method := reflect.ValueOf(impl).MethodByName(methodName)
		in := make([]reflect.Value, method.Type().NumIn())
		var handlingClosure = func(w http.ResponseWriter, r *http.Request) {
			s := unboxSingleString(w, r)
			in[0] = reflect.ValueOf(s)
			method.Call(in)
		}

		err := api.AddHandler(methodName, "POST", "/"+methodName, methodName, handlingClosure)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Registered " + methodName)
		}
	}
	return
}

func unboxSingleString(w http.ResponseWriter, r *http.Request) string {

	fmt.Println("Someone called print")
	// dont need to cache ?
	w.Header().Set("Cache-Control", "no-store")

	r.ParseForm()           // Parses the request body
	x := r.Form.Get("data") // x will be "" if parameter is not set

	qs := rest.ParseQueryStrings(r)
	message := fmt.Sprintf("Welcome to restfool take a look at https://%s/help", r.Host)
	msg := rest.Msg{Message: message}
	rest.EncodeAndSend(w, r, qs, msg)
	return x
}
