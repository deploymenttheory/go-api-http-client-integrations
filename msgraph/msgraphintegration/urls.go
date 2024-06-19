// apiintegrations/msgraph/urls.go
package msgraphintegration

// GetBaseDomain returns the base domain for the Jamf Pro integration.
func (m *Integration) GetBaseDomain() string {
	return m.BaseDomain
}
