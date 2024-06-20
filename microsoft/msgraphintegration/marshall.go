// microsoft/msgraphintegration/marshall.go
package msgraphintegration

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/deploymenttheory/go-api-http-client-integrations/shared/helpers"
	"go.uber.org/zap"
)

// methodsWithBody is a set of HTTP methods that benefit from logging the request body.
var methodsWithBody = map[string]bool{
	"POST":  true,
	"PUT":   true,
	"PATCH": true,
}

// MarshalRequest encodes the request body as JSON for the Microsoft Graph API.
// This function takes an interface{} type body, an HTTP method, and an endpoint as input,
// and returns the marshaled JSON byte slice along with any error encountered during marshaling.
// The function ensures that the request body is always marshaled as JSON.
// It logs the JSON request body for POST, PUT, and PATCH methods using the integrated logger.
//
// Parameters:
//   - body: The request body to be marshaled, of type interface{}.
//   - method: The HTTP method being used for the request (e.g., "POST", "PUT", "PATCH").
//   - endpoint: The API endpoint for the request.
//
// Returns:
//   - []byte: The marshaled JSON byte slice of the request body.
//   - error: Any error encountered during the marshaling process.
//
// Logging:
//   - Logs an error if JSON marshaling fails.
//   - Logs the JSON request body for POST, PUT, and PATCH methods.
//
// Set of methods that require logging the request body
func (m *Integration) marshalRequest(body interface{}, method string, endpoint string) ([]byte, error) {
	var (
		data []byte
		err  error
	)

	// Marshal the body as JSON
	data, err = json.Marshal(body)
	if err != nil {
		m.Logger.Error("Failed marshaling JSON request", zap.Error(err))
		return nil, err
	}

	// Log the JSON request body for methods that require it
	if methodsWithBody[method] {
		m.Logger.Debug("JSON Request Body", zap.String("Body", string(data)), zap.String("Endpoint", endpoint))
	} else {
		m.Logger.Debug("Request Endpoint", zap.String("Endpoint", endpoint))
	}

	return data, nil
}

// MarshalMultipartRequest handles multipart form data encoding with secure file handling and returns the encoded body and content type.
func (m *Integration) marshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for field, value := range fields {
		if err := writer.WriteField(field, value); err != nil {
			return nil, "", err
		}
	}

	for formField, filePath := range files {
		file, err := helpers.SafeOpenFile(filePath)
		if err != nil {
			m.Logger.Error("Failed to open file securely", zap.String("file", filePath), zap.Error(err))
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

	contentType := writer.FormDataContentType()
	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return body.Bytes(), contentType, nil
}
