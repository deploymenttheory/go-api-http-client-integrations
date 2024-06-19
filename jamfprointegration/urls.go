// jamfpro_api_url.go
package jamfprointegration

// SetBaseDomain returns the appropriate base domain for URL construction.
// It uses j.OverrideBaseDomain if set, otherwise falls back to DefaultBaseDomain.
func (j *Integration) GetBaseDomain() string {
	return j.BaseDomain
}
