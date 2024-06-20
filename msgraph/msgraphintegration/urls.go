// apiintegrations/msgraph/urls.go
package msgraphintegration

// GetTenantID returns the tenant ID for the Microsoft Graph integration.
func (m *Integration) GetTenantID() string {
	return m.TenantID
}

// getFQDN returns the fully qualified domain name for Microsoft Graph.
func (m *Integration) getFQDN() string {
	return "https://graph.microsoft.com"
}
