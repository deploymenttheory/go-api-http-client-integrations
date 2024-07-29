package jamfprointegration

import (
	"strings"

	"go.uber.org/zap"
)

const WeightedAcceptHeader = "application/x-x509-ca-cert;q=0.95," +
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
	"*/*;q=0.05"

// getContentTypeHeader determines the appropriate Content-Type header for a given API endpoint.
// It sets the Content-Type to "application/octet-stream" specifically for the endpoint "/api/v1/packages/{id}/upload".
// For other endpoints, it attempts to match the Content-Type based on the endpoint pattern:
// - For URL endpoints starting with "/JSSResource", it defaults to "application/xml" for the Classic API.
// - For URL endpoints starting with "/api", it defaults to "application/json" for the JamfPro API.
// If the endpoint does not match any of the predefined patterns, "application/json" is used as a fallback.
// This method logs the decision process at various stages for debugging purposes.
func (j *Integration) getContentTypeHeader(endpoint string) string {
	j.Sugar.Debug("Determining Content-Type for endpoint", zap.String("endpoint", endpoint))

	// TODO change this contains to regex. We want to rule out malformed endpoints with multiple occurances.
	if strings.Contains(endpoint, "/api/v1/packages/") && strings.Contains(endpoint, "/upload") {
		j.Sugar.Debugw("Content-Type for packages upload endpoint set to application/octet-stream", "endpoint", endpoint)
		return "application/octet-stream"
	}

	if strings.Contains(endpoint, "/JSSResource") {
		j.Sugar.Debugw("Content-Type for endpoint defaulting to XML for Classic API", "endpoint", endpoint)
		// TODO should this be application/xml or text/xml?
		return "application/xml"
	}

	if strings.Contains(endpoint, "/api") {
		j.Sugar.Debugw("Content-Type for endpoint defaulting to JSON for JamfPro API", "endpoint", endpoint)
		return "application/json"
	}

	// j.Sugar.Warnw("Content-Type for endpoint not found in configMap or standard patterns, using default JSON", "endpoint", endpoint)
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
	return WeightedAcceptHeader
}

// getUserAgentHeader returns the User-Agent header string for the Jamf Pro API.
func (j *Integration) getUserAgentHeader() string {
	return "go-api-http-client-jamfpro-integration"
}
