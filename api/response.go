package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response from the Neocities API
type Response struct {
	Result    string `json:"result"`
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
	Info      Info
	Body      []byte
}

// Info contains the properties for a site
type Info struct {
	Sitename    string   `json:"sitename"`
	Hits        int32    `json:"hits"`
	CreatedAt   string   `json:"created_at"`
	LastUpdated string   `json:"last_updated"`
	Domain      string   `json:"domain"`
	Tags        []string `json:"tags"`
}

// PopulateFromHTTPResponse use a HTTP response to populate itself
func (r *Response) PopulateFromHTTPResponse(res *http.Response) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.Unmarshal(body, &r)

	r.Body = body

	return err
}

// Print is printing the contents of the response to stdout
func (r *Response) Print() {
	fmt.Println("Result:   ", r.Result)

	if r.ErrorType != "" {
		fmt.Println("ErrorType:", r.ErrorType)
	}

	fmt.Println("Message:  ", r.Message)
}
