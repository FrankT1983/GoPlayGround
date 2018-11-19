package impls

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	help(toPrint)
}

func help(method string) ([]byte, error) {
	fmt.Println("start building request")
	e := eater.NewEater("127.0.0.1", 9443)
	e.SetBasicAuth("testuser", "testpass")
	e.SetVerifyTLS(false)

	req := e.CreateRequest("/Print", "POST", nil)

	fmt.Println("send request")
	resp, err := req.Go()
	if err != nil {

		fmt.Println("ups ... error " + string(err.Error()))
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {

		fmt.Println("status OK")
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if strings.Contains(string(b), "error") {

		}
		return b, nil
	}

	fmt.Println("status not OK")
	fmt.Printf("statusCode: %d - %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	return nil, fmt.Errorf("statusCode: %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}
