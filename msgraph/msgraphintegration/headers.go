// msgraph/msgraphintegration/headers.go
package msgraphintegration

import (
	"strings"

	"github.com/deploymenttheory/go-api-http-client/logger"
	"go.uber.org/zap"
)

// getContentTypeHeader determines the appropriate Content-Type header for a given API endpoint.
// It attempts to find a content type that matches the endpoint prefix in the global configMap.
// If a match is found and the content type is defined (not nil), it returns the specified content type.
// If the endpoint does not match any of the predefined patterns, "application/json" is used as a fallback.
// This method logs the decision process at various stages for debugging purposes.
func (m *Integration) getContentTypeHeader(endpoint string, log logger.Logger) string {
	// Dynamic lookup from configuration should be the first priority
	for key, config := range configMap {
		if strings.HasPrefix(endpoint, key) {
			if config.ContentType != nil {
				m.Logger.Debug("Content-Type for endpoint found in configMap", zap.String("endpoint", endpoint), zap.String("content_type", *config.ContentType))
				return *config.ContentType
			}
			m.Logger.Debug("Content-Type for endpoint is nil in configMap, handling as special case", zap.String("endpoint", endpoint))
			// If a nil ContentType is an expected case, do not set Content-Type header.
			return "" // Return empty to indicate no Content-Type should be set.
		}
	}

	// Fallback to JSON if no other match is found.
	m.Logger.Debug("Content-Type for endpoint not found in configMap or standard patterns, using default JSON", zap.String("endpoint", endpoint))
	return "application/json"
}

// GetAcceptHeader constructs and returns a weighted Accept header string for HTTP requests.
// The Accept header indicates the MIME types that the client can process and prioritizes them
// based on the quality factor (q) parameter. Higher q-values signal greater preference.
// This function specifies a range of MIME types with their respective weights, ensuring that
// the server is informed of the client's versatile content handling capabilities while
// indicating a preference for XML. The specified MIME types cover common content formats like
// images, JSON, XML, HTML, plain text, and certificates, with a fallback option for all other types.
func (m *Integration) getAcceptHeader() string {
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

// getUserAgentHeader returns the User-Agent header string for the Microsoft Graph API.
func (m *Integration) getUserAgentHeader() string {
	return "go-api-http-client-msgraph-integration"
}
