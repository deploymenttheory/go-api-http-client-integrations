// jamfpro_api_request.go
package jamfprointegration

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

// MarshalRequest encodes the request body according to the endpoint for the API.
func (j *Integration) marshalRequest(body interface{}, method string, endpoint string) ([]byte, error) {
	log.Println("start of marshal")
	var (
		data []byte
		err  error
	)

	log.Println("marshal flag 2")
	// Determine the format based on the endpoint
	format := "json"
	if strings.Contains(endpoint, "/JSSResource") {
		format = "xml"
	} else if strings.Contains(endpoint, "/api") {
		format = "json"
	}

	log.Println("marshal flag 3")

	switch format {
	case "xml":
		log.Println("marshal flag 4")
		data, err = xml.Marshal(body)
		if err != nil {
			return nil, err
		}
		log.Println("marshal flag 5")

		if method == "POST" || method == "PUT" {
			j.Logger.Debug("XML Request Body", zap.String("Body", string(data)))
		}
		log.Println("marshal flag 6")
		return data, nil

	case "json":
		data, err = json.Marshal(body)
		if err != nil {
			j.Logger.Error("Failed marshaling JSON request", zap.Error(err))
			return nil, err
		}

		if method == "POST" || method == "PUT" || method == "PATCH" {
			j.Logger.Debug("JSON Request Body", zap.String("Body", string(data)))
		}

		return data, nil

	default:
		return nil, errors.New("invalid marshal format")
	}
}

// MarshalMultipartRequest handles multipart form data encoding with secure file handling and returns the encoded body and content type.
func (j *Integration) marshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the simple fields to the form data
	for field, value := range fields {
		if err := writer.WriteField(field, value); err != nil {
			return nil, "", err
		}
	}

	// Add the files to the form data, using safeOpenFile to ensure secure file access
	for formField, filePath := range files {
		file, err := SafeOpenFile(filePath)
		if err != nil {
			j.Logger.Error("Failed to open file securely", zap.String("file", filePath), zap.Error(err))
			return nil, "", err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(formField, filepath.Base(filePath))
		if err != nil {
			return nil, "", err
		}
		if _, err := io.Copy(part, file); err != nil {
			return nil, "", err
		}
	}

	// Close the writer to finish writing the multipart message
	contentType := writer.FormDataContentType()
	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return body.Bytes(), contentType, nil
}
