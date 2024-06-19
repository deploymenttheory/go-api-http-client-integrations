package msgraphintegration

// Endpoint constants represent the URL suffixes used for Jamf API token interactions.
const (
	// Auth
	oAuthTokenEndpoint      string = "/oauth2/v2.0/token"
	bearerTokenEndpoint     string = "graph.microsoft.com"
	invalidateTokenEndpoint string = "graph.microsoft.com"
	oAuthTokenScope         string = "https://graph.microsoft.com/.default"
	baseAuthURL             string = "https://login.microsoftonline.com"
)
