package jamfprointegration

import "time"

// Endpoint constants represent the URL suffixes used for Jamf API token interactions.
const (
	// Auth
	oAuthTokenEndpoint      string = "/api/v1/oauth/token"
	bearerTokenEndpoint     string = "/api/v1/auth/token"
	invalidateTokenEndpoint string = "/api/v1/auth/invalidate-token"
	keepAliveTokenEndpoint  string = "/api/v1/auth/keep-alive"

	// Load balancer workaround
	LoadBalancerTargetCookie string        = "jpro-ingress"
	LoadBalancerPollCount    int           = 5
	LoadBalancerTimeOut      time.Duration = 7 * time.Second
)
