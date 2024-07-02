package jamfprointegration

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/deploymenttheory/go-api-http-client-integrations/shared/helpers"
	"go.uber.org/zap"
)

// MarshalRequest encodes the request body according to the endpoint for the Jamf Pro API.
// This function marshals the request body as JSON or XML based on the endpoint string.
// It takes an interface{} type body, an HTTP method, and an endpoint as input, and returns the
// marshaled byte slice along with any error encountered during marshaling.
//
// Parameters:
//   - body: The request body to be marshaled, of type interface{}.
//   - method: The HTTP method being used for the request (e.g., "POST", "PUT", "PATCH").
//   - endpoint: The API endpoint for the request.
//
// Returns:
//   - []byte: The marshaled byte slice of the request body.
//   - error: Any error encountered during the marshaling process.
//
// Functionality:
//   - Determines the format (JSON or XML) based on the endpoint string.
//   - Marshals the body as JSON if the endpoint contains "/api" or as XML if it contains "/JSSResource".
//   - Logs the marshaled request body for POST, PUT, and PATCH methods using the integrated logger.
//   - Logs an error if marshaling fails and returns the error.
//   - Returns an error if the format is invalid.
func (j *Integration) marshalRequest(body interface{}, method string, endpoint string) ([]byte, error) {
	var (
		data []byte
		err  error
	)

	// Determine the format based on the endpoint
	format := "json"
	if strings.Contains(endpoint, "/JSSResource") {
		format = "xml"
	} else if strings.Contains(endpoint, "/api") {
		format = "json"
	}

	switch format {
	case "xml":
		data, err = xml.Marshal(body)
		if err != nil {
			return nil, err
		}

		if method == "POST" || method == "PUT" {
			j.Sugar.Debug("XML Request Body", zap.String("Body", string(data)))
		}

		return data, nil

	case "json":
		data, err = json.Marshal(body)
		if err != nil {
			j.Sugar.Error("Failed marshaling JSON request", zap.Error(err))
			return nil, err
		}

		if method == "POST" || method == "PUT" || method == "PATCH" {
			// TODO it hates this, pointer dereference on this log? Weird.
			j.Sugar.Debug("JSON Request Body:", zap.Any("body", json.RawMessage(data)))

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

	for field, value := range fields {
		if err := writer.WriteField(field, value); err != nil {
			return nil, "", err
		}
	}

	for formField, filePath := range files {
		file, err := helpers.SafeOpenFile(filePath)
		if err != nil {
			j.Sugar.Error("Failed to open file securely", zap.String("file", filePath), zap.Error(err))
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
