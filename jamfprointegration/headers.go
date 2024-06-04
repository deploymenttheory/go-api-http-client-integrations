package jamfprointegration

import (
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

func (j *Integration) setRequestHeaders(req *http.Request) error {
	req.Header.Add("Accept", j.getAcceptHeader())
	req.Header.Add("Content-Type", j.getContentTypeHeader(req.URL.String()))
	req.Header.Add("User-Agent", j.getUserAgentHeader())

	token, err := j.Token()
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return nil
}

// GetContentTypeHeader determines the appropriate Content-Type header for a given API endpoint.
// It attempts to find a content type that matches the endpoint prefix in the global configMap.
// If a match is found and the content type is defined (not nil), it returns the specified content type.
// If the content type is nil or no match is found in configMap, it falls back to default behaviors:
// - For url endpoints starting with "/JSSResource", it defaults to "application/xml" for the Classic API.
// - For url endpoints starting with "/api", it defaults to "application/json" for the JamfPro API.
// If the endpoint does not match any of the predefined patterns, "application/json" is used as a fallback.
// This method logs the decision process at various stages for debugging purposes.
func (j *Integration) getContentTypeHeader(endpoint string) string {
	if strings.Contains(endpoint, "/JSSResource") {
		j.Logger.Debug("Content-Type for endpoint defaulting to XML for Classic API", zap.String("endpoint", endpoint))
		return "application/xml"
	} else if strings.Contains(endpoint, "/api") {
		j.Logger.Debug("Content-Type for endpoint defaulting to JSON for JamfPro API", zap.String("endpoint", endpoint))
		return "application/json"
	}

	j.Logger.Debug("Content-Type for endpoint not found in configMap or standard patterns, using default JSON", zap.String("endpoint", endpoint))
	return "application/json"
}

// GetAcceptHeader constructs and returns a weighted Accept header string for HTTP requests.
// The Accept header indicates the MIME types that the client can process and prioritizes them
// based on the quality factor (q) parameter. Higher q-values signal greater preference.
// This function specifies a range of MIME types with their respective weights, ensuring that
// the server is informed of the client's versatile content handling capabilities while
// indicating a preference for XML. The specified MIME types cover common content formats like
// images, JSON, XML, HTML, plain text, and certificates, with a fallback option for all other types.
func (j *Integration) getAcceptHeader() string {
	weightedAcceptHeader := "application/x-x509-ca-cert;q=0.95," +
		"application/pkix-cert;q=0.94," +
		"application/pem-certificate-chain;q=0.93," +
		"application/octet-stream;q=0.8," + // For general binary files
		"image/png;q=0.75," +
		"image/jpeg;q=0.74," +
		"image/*;q=0.7," +
		"application/xml;q=0.65," +
		"text/xml;q=0.64," +
		"text/xml;charset=UTF-8;q=0.63," +
		"application/json;q=0.5," +
		"text/html;q=0.5," +
		"text/plain;q=0.4," +
		"*/*;q=0.05" // Fallback for any other types

	return weightedAcceptHeader
}

func (j *Integration) getUserAgentHeader() string {
	return "go-api-http-client-jamfpro-integration"
}
