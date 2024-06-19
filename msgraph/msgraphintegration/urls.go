// apiintegrations/msgraph/urls.go
package msgraphintegration

// GetTenantName returns the tenant name for the Microsoft Graph integration.
func (m *Integration) GetTenantName() string {
	return m.TenantName
}
