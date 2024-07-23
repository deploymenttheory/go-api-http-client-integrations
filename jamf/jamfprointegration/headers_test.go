package jamfprointegration

import (
	"testing"
)

func TestIntegration_getContentTypeHeader(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name        string
		integration *Integration
		args        args
		want        string
	}{
		{
			name:        "package endpoint",
			integration: test_MockIntegration(),
			args: args{
				endpoint: "/api/v1/packages/upload",
			},
			want: "application/octet-stream",
		},
		{
			name:        "classic endpoint",
			integration: test_MockIntegration(),
			args: args{
				endpoint: "/JSSResource/a_resource",
			},
			want: "application/xml",
		},
		{
			name:        "pro endpoint",
			integration: test_MockIntegration(),
			args: args{
				endpoint: "/api/v1/a_resource",
			},
			want: "application/json",
		},
		{
			name:        "unexpected endpoint",
			integration: test_MockIntegration(),
			args: args{
				endpoint: "/not/an/endpoint",
			},
			want: "application/json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.integration
			if got := j.getContentTypeHeader(tt.args.endpoint); got != tt.want {
				t.Errorf("Integration.getContentTypeHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegration_getAcceptHeader(t *testing.T) {
	tests := []struct {
		name        string
		integration *Integration
		want        string
	}{
		{
			name:        "1",
			integration: test_MockIntegration(),
			want: "application/x-x509-ca-cert;q=0.95," +
				"application/pkix-cert;q=0.94," +
				"application/pem-certificate-chain;q=0.93," +
				"application/octet-stream;q=0.8," +
				"image/png;q=0.75," +
				"image/jpeg;q=0.74," +
				"image/*;q=0.7," +
				"application/xml;q=0.65," +
				"text/xml;q=0.64," +
				"text/xml;charset=UTF-8;q=0.63," +
				"application/json;q=0.5," +
				"text/html;q=0.5," +
				"text/plain;q=0.4," +
				"*/*;q=0.05",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.integration
			if got := j.getAcceptHeader(); got != tt.want {
				t.Errorf("Integration.getAcceptHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegration_getUserAgentHeader(t *testing.T) {
	tests := []struct {
		name        string
		integration *Integration
		want        string
	}{
		{
			name:        "1",
			integration: test_MockIntegration(),
			want:        "go-api-http-client-jamfpro-integration",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.integration
			if got := j.getUserAgentHeader(); got != tt.want {
				t.Errorf("Integration.getUserAgentHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
