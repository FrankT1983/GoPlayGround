package resthelper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	eater "github.com/cseeger-epages/resteater-go"
	rest "github.com/cseeger-epages/restfool-go"
)

// AddInterfaceHandlers generate the handlers for a interface ... code generatio would be the way to go here
func AddInterfaceHandlers(api rest.RestAPI, interfaceType reflect.Type, impl interface{}) {

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

// CallStringFunctionOferRest calls a function over rest that requires only a single string input
func CallStringFunctionOferRest(functionName string, toSend string) {
	fmt.Println("start building request")
	e := eater.NewEater("127.0.0.1", 9443)
	e.SetBasicAuth("testuser", "testpass")
	e.SetVerifyTLS(false)

	req := e.CreateRequest("/"+functionName, "POST", url.Values{"data": {toSend}})
	resp, err := req.Go()
	if err != nil {

		fmt.Println("ups ... error " + string(err.Error()))
		return
	}

	if resp.StatusCode == http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		if strings.Contains(string(b), "error") {

		}
		return
	}

	fmt.Printf("statusCode: %d - %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	return
}
