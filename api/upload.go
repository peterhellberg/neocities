package api

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

// UploadData contains the filename and content
type UploadData struct {
	FileName string
	Content  []byte
}

func Upload(a Authenticator, data []UploadData) (Response, error) {
	req, err := newUploadDataRequest(a, data)
	if err != nil {
		return Response{}, err
	}

	return performHTTPRequest(req)
}

// Create a new upload data request
func newUploadDataRequest(a Authenticator, data []UploadData) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the contents of each file to the multipart body
	for _, d := range data {
		part, err := writer.CreateFormFile(d.FileName, d.FileName)
		if err != nil {
			return nil, err
		}

		if _, err = part.Write(d.Content); err != nil {
			return nil, err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL+"upload", body)
	if err != nil {
		return req, err
	}

	if a != nil {
		a.Authenticate(req)
	}

	// Set the content type of the form data
	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req, nil
}
