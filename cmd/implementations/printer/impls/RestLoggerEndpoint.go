package impls

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	eater "github.com/cseeger-epages/resteater-go"
)

// RestLoggerEndpoint implemnet the IPrinter interface by providing an REST endpoint.
type RestLoggerEndpoint struct {
	LoggerName string
}

// PrintToLog implements the IPrinter interface function.
func (l RestLoggerEndpoint) PrintToLog(toPrint string) {
	// send rest request
	callStringFunctionOferRest("PrintToLog", toPrint)
}

func callStringFunctionOferRest(functionName string, toSend string) {
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
