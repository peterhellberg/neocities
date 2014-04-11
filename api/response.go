package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/peterhellberg/neocities/utils"
)

// Response from the Neocities API
type Response struct {
	Result    string `json:"result"`
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
}

// PopulateFromHTTPResponse use a HTTP response to populate itself
func (r *Response) PopulateFromHTTPResponse(res *http.Response) {
	body, err := ioutil.ReadAll(res.Body)
	utils.Check(err)

	err = json.Unmarshal(body, &r)
	utils.Check(err)
}

// Print is printing the contents of the response to stdout
func (r *Response) Print() {
	fmt.Println("Result:   ", r.Result)
	if r.ErrorType != "" {
		fmt.Println("ErrorType:", r.ErrorType)
	}
	fmt.Println("Message:  ", r.Message)
}
