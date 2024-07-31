package test

import (
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-api-http-client-integrations/jamf/jamfprointegration"
)

func TestIntegration_getSessionCookies(t *testing.T) {
	type fields struct {
		Integration *jamfprointegration.Integration
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*http.Cookie
		wantErr bool
	}{
		{
			name: "get session cookies ok",
			fields: fields{
				Integration: NewIntegrationFromEnv(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.fields.Integration
			_, err := j.GetSessionCookies()
			if (err != nil) != tt.wantErr {
				t.Errorf("Integration.getSessionCookies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
