package jamfprointegration

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"go.uber.org/zap"
)

type fixedPlatformAuth struct{}

func (fixedPlatformAuth) getNewToken() error       { return errors.New("unexpected token request") }
func (fixedPlatformAuth) getTokenString() string   { return "test-token" }
func (fixedPlatformAuth) getExpiryTime() time.Time { return time.Now().Add(time.Hour) }
func (fixedPlatformAuth) tokenExpired() bool       { return false }
func (fixedPlatformAuth) tokenInBuffer() bool      { return false }
func (fixedPlatformAuth) tokenEmpty() bool         { return false }

func TestPlatformGatewayConstructURLUsesHeaderScope(t *testing.T) {
	integration := &Integration{
		JamfProFQDN:          "https://gateway.example",
		AuthMethodDescriptor: "platform",
		TenantID:             "tenant-id",
	}

	tests := []struct {
		endpoint string
		want     string
	}{
		{endpoint: "/JSSResource/computers", want: "https://gateway.example/api/proclassic/computers"},
		{endpoint: "/api/v1/computers", want: "https://gateway.example/api/pro/v1/computers"},
		{endpoint: "/api/v1", want: "https://gateway.example/api/pro/v1"},
	}

	for _, tt := range tests {
		t.Run(tt.endpoint, func(t *testing.T) {
			if got := integration.ConstructURL(tt.endpoint); got != tt.want {
				t.Fatalf("ConstructURL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPlatformGatewayPrepRequestAddsScopeHeader(t *testing.T) {
	tests := []struct {
		name              string
		tenantID          string
		environmentID     string
		wantTenantID      string
		wantEnvironmentID string
	}{
		{name: "tenant", tenantID: "tenant-id", wantTenantID: "tenant-id"},
		{name: "environment", environmentID: "environment-id", wantEnvironmentID: "environment-id"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			integration := &Integration{
				AuthMethodDescriptor: "platform",
				Sugar:                zap.NewNop().Sugar(),
				auth:                 fixedPlatformAuth{},
				TenantID:             tt.tenantID,
				EnvironmentID:        tt.environmentID,
			}
			req, err := http.NewRequest(http.MethodGet, "https://gateway.example/api/pro/v1/computers", nil)
			if err != nil {
				t.Fatal(err)
			}

			if err := integration.PrepRequestParamsAndAuth(req); err != nil {
				t.Fatal(err)
			}
			if got := req.Header.Get("X-Tenant-Id"); got != tt.wantTenantID {
				t.Errorf("X-Tenant-Id = %q, want %q", got, tt.wantTenantID)
			}
			if got := req.Header.Get("X-Environment-Id"); got != tt.wantEnvironmentID {
				t.Errorf("X-Environment-Id = %q, want %q", got, tt.wantEnvironmentID)
			}
		})
	}
}

func TestPlatformGatewayPrepRequestRejectsAmbiguousScope(t *testing.T) {
	for _, tt := range []struct {
		name          string
		tenantID      string
		environmentID string
	}{
		{name: "missing"},
		{name: "both", tenantID: "tenant-id", environmentID: "environment-id"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			integration := &Integration{
				AuthMethodDescriptor: "platform",
				Sugar:                zap.NewNop().Sugar(),
				auth:                 fixedPlatformAuth{},
				TenantID:             tt.tenantID,
				EnvironmentID:        tt.environmentID,
			}
			req, err := http.NewRequest(http.MethodGet, "https://gateway.example/api/pro/v1/computers", nil)
			if err != nil {
				t.Fatal(err)
			}
			if err := integration.PrepRequestParamsAndAuth(req); err == nil {
				t.Fatal("PrepRequestParamsAndAuth() succeeded with an ambiguous scope")
			}
		})
	}
}

func TestBuildWithPlatformGatewayOAuthWithScopeRejectsAmbiguousScope(t *testing.T) {
	for _, tt := range []struct {
		name          string
		tenantID      string
		environmentID string
	}{
		{name: "missing"},
		{name: "both", tenantID: "tenant-id", environmentID: "environment-id"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := BuildWithPlatformGatewayOAuthWithScope(
				"https://gateway.example",
				zap.NewNop().Sugar(),
				time.Minute,
				"client-id",
				"client-secret",
				tt.tenantID,
				tt.environmentID,
				true,
				http.Client{},
			)
			if err == nil {
				t.Fatal("BuildWithPlatformGatewayOAuthWithScope() succeeded with an ambiguous scope")
			}
		})
	}
}
