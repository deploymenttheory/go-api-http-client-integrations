package jamfprointegration

import (
	"testing"
)

func TestIntegration_getContentTypeHeader(t *testing.T) {
	testIntegration := newIntegrationWithLogger()
	type fields struct {
		integration Integration
	}
	type args struct {
		endpoint string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "packages endpoint",
			fields: fields{
				integration: testIntegration,
			},
			args: args{
				endpoint: "http://yourserver.jamfcloud.com/api/v1/packages/upload/something else",
			},
			want: "application/octet-stream",
		},
		{
			name: "classic endpoint",
			fields: fields{
				integration: testIntegration,
			},
			args: args{
				endpoint: "http://yourserver.jamfcloud.com/JSSResource/somethingelse",
			},
			want: "application/xml",
		},
		{
			name: "pro endpoint",
			fields: fields{
				integration: testIntegration,
			},
			args: args{
				endpoint: "http://yourserver.jamfcloud.com/api/v100/cheeseburger",
			},
			want: "application/json",
		},
		{
			name: "unrecognised endpoint",
			fields: fields{
				integration: testIntegration,
			},
			args: args{
				endpoint: "http://yourserver.jamfcloud.com/spaghetti",
			},
			want: "application/json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.fields.integration
			if got := j.getContentTypeHeader(tt.args.endpoint, ""); got != tt.want {
				t.Errorf("Integration.getContentTypeHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegration_getAcceptHeader(t *testing.T) {
	type fields struct {
		integration Integration
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "singular",
			fields: fields{
				integration: Integration{},
			},
			want: WeightedAcceptHeader,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.fields.integration
			if got := j.getAcceptHeader(); got != tt.want {
				t.Errorf("Integration.getAcceptHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegration_getUserAgentHeader(t *testing.T) {
	type fields struct {
		integration Integration
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "singular",
			fields: fields{
				integration: Integration{},
			},
			want: UserAgentHeader,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.fields.integration
			if got := j.getUserAgentHeader(); got != tt.want {
				t.Errorf("Integration.getUserAgentHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
