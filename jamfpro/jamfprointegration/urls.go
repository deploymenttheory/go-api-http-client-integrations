// jamfpro/jamfprointegration/urls.go
package jamfprointegration

// GetBaseDomain returns the base domain for the Jamf Pro integration.
func (j *Integration) GetBaseDomain() string {
	return j.BaseDomain
}
