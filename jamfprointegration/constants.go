package jamfprointegration

// Endpoint constants represent the URL suffixes used for Jamf API token interactions.
const (
	oAuthTokenEndpoint      string = "/api/oauth/token"
	bearerTokenEndpoint     string = "/api/v1/auth/token"
	invalidateTokenEndpoint string = "/api/v1/auth/invalidate-token"
	keepAliveTokenEndpoint  string = "/api/v1/auth/keep-alive"
)

type auth interface {
	tokenExpired() bool
	tokenInBuffer() bool
	tokenEmpty() bool
	getNewToken() error
}
