package resthelper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strings"

	"encoding/json"

	eater "github.com/cseeger-epages/resteater-go"
	rest "github.com/cseeger-epages/restfool-go"
)

// ServerInfo encapsules simpe authentication information for a server
type ServerInfo struct {
	Server   string
	Port     int
	User     string
	Password string
}

// AddInterfaceHandlers generate the handlers for a interface ... code generatio would be the way to go here
func AddInterfaceHandlers(api rest.RestAPI, interfaceType reflect.Type, impl interface{}) {

	for i := 0; i < interfaceType.NumMethod(); i++ {
		methodName := interfaceType.Method(i).Name

		method := reflect.ValueOf(impl).MethodByName(methodName)
		if method.Type().NumIn() != 1 {
			log.Fatal("Cannot jet automatically handle more than 1 parameter for function " + methodName)
			continue
		}

		in := make([]reflect.Value, method.Type().NumIn())

		var handlingClosure = func(w http.ResponseWriter, r *http.Request) {
		}

		// todo: use a switch?
		if method.Type().In(0).Kind() == reflect.String {
			handlingClosure = func(w http.ResponseWriter, r *http.Request) {
				s := unboxSingleString(w, r)
				in[0] = reflect.ValueOf(s)
				method.Call(in)
			}
		} else if method.Type().In(0).Kind() == reflect.Struct {
			handlingClosure = func(w http.ResponseWriter, r *http.Request) {
				inS := reflect.New(method.Type().In(0))
				unboxSingleStruct(w, r, &inS)
				in[0] = reflect.ValueOf(inS.Elem().Interface())
				method.Call(in)
			}
		} else {
			log.Fatal("Input parameter type not defined for " + methodName + " : " + method.Type().In(0).Name())
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

func unboxSingleStruct(w http.ResponseWriter, r *http.Request, valueToFill *reflect.Value) {
	// dont need to cache ?
	w.Header().Set("Cache-Control", "no-store")

	r.ParseForm()           // Parses the request body
	x := r.Form.Get("data") // x will be "" if parameter is not set

	structToFill := valueToFill.Interface()
	if err := json.Unmarshal([]byte(x), &structToFill); err != nil {
		fmt.Println("Failure during unmarshal")
		panic(err)
	}

	qs := rest.ParseQueryStrings(r)
	message := fmt.Sprintf("Welcome to restfool take a look at https://%s/help", r.Host)
	msg := rest.Msg{Message: message}
	rest.EncodeAndSend(w, r, qs, msg)
	return
}

func createRequest(info ServerInfo, functionName string, form url.Values) eater.Request {
	e := eater.NewEater(info.Server, info.Port)
	e.SetBasicAuth(info.User, info.Password)
	e.SetVerifyTLS(false)
	req := e.CreateRequest("/"+functionName, "POST", form)
	return req
}

func handleResponseAndErrors(resp *http.Response, err error) {
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

}

// CallStringFunctionOverRest calls a function over rest that requires only a single string input
func CallStringFunctionOverRest(info ServerInfo, functionName string, toSend string) {
	req := createRequest(info, functionName, url.Values{"data": {toSend}, "type": {"string"}})
	resp, err := req.Go()
	handleResponseAndErrors(resp, err)
	return
}

//CallStructFunctionOverRest calls a function over rest that requires only a single string input
func CallStructFunctionOverRest(info ServerInfo, functionName string, toSend interface{}) {

	b, err := json.Marshal(toSend)
	if err != nil {
		log.Fatal(err)
		return
	}

	req := createRequest(info, functionName, url.Values{"data": {string(b)}, "type": {"json"}})

	resp, err := req.Go()
	handleResponseAndErrors(resp, err)
	return
}

// WhereAmI returns the simple name of the function in which it is called
// got this from: https://lawlessguy.wordpress.com/2016/04/17/display-file-function-and-line-number-in-go-golang/
func WhereAmI(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, _, _, _ := runtime.Caller(depth)
	fullName := runtime.FuncForPC(function).Name()
	tmp := strings.Split(fullName, ".")
	return tmp[len(tmp)-1]
}
