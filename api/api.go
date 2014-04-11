package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/peterhellberg/neocities/utils"
)

const (
	apiURL = "https://neocities.org/api/"
)

// Credentials contains the username and password
type Credentials struct {
	User string
	Pass string
}

// UploadFiles takes a set of credentials and
// a list of filename paths to upload to Neocities.
func UploadFiles(cred *Credentials, paths []string) (Response, error) {
	req, err := prepareUploadPOSTRequest(cred, paths)
	utils.Check(err)

	res, err := sendHTTPRequest(req)
	utils.Check(err)

	defer res.Body.Close()

	var r Response

	r.PopulateFromHTTPResponse(res)

	if res.StatusCode == 200 {
		return r, nil
	}

	return r, utils.Error("unsuccessful")
}

// Prepare upload POST request
func prepareUploadPOSTRequest(cred *Credentials, paths []string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the contents of each file to the multipart body
	for _, path := range paths {
		file, err := os.Open(path)

		defer file.Close()

		if err != nil {
			return nil, err
		}

		base := filepath.Base(path)
		part, err := writer.CreateFormFile(base, base)

		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, file)
		utils.Check(err)
	}

	err := writer.Close()
	utils.Check(err)

	req, err := http.NewRequest("POST", apiURL+"upload", body)
	if err != nil {
		return req, err
	}

	// Authenticate using the user and pass
	req.SetBasicAuth(cred.User, cred.Pass)

	// Set the content type of the form data
	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req, nil
}

// Send HTTP request
func sendHTTPRequest(req *http.Request) (*http.Response, error) {
	// Create a HTTP client
	client := &http.Client{}

	// Perform the POST request
	res, err := client.Do(req)
	utils.Check(err)

	return res, nil
}
