package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFiles takes a set of credentials and
// a list of filename paths to upload to Neocities.
func UploadFiles(a Authenticator, paths []string) (Response, error) {
	req, err := newUploadRequest(a, paths)
	if err != nil {
		return Response{}, err
	}

	return performHTTPRequest(req)
}

// Create a new upload request
func newUploadRequest(a Authenticator, paths []string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the contents of each file to the multipart body
	for _, path := range paths {
		filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			file, err := os.Open(p)

			defer file.Close()

			if err != nil {
				return err
			}

			part, err := writer.CreateFormFile(p, p)

			if err != nil {
				return err
			}

			if _, err := io.Copy(part, file); err != nil {
				return err
			}

			return nil
		})
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL+"upload", body)
	if err != nil {
		return nil, err
	}

	if a != nil {
		a.Authenticate(req)
	}

	// Set the content type of the form data
	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req, nil
}
