package jamfprointegration

import (
	"net/http"
	"reflect"
	"testing"

	"go.uber.org/zap"
)

type xmlTestBody struct {
	Name string `xml:"name"`
}

func TestIntegration_marshalRequest(t *testing.T) {
	type fields struct {
		jamfProFQDN          string
		AuthMethodDescriptor string
		Sugar                *zap.SugaredLogger
		auth                 authInterface
		http                 http.Client
	}
	type args struct {
		body     any
		method   string
		endpoint string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "GET without body sends no bytes",
			args: args{body: nil, method: http.MethodGet, endpoint: "/api/v1/computers"},
			want: nil,
		},
		{
			name: "GET ignores a supplied body",
			args: args{body: struct{}{}, method: http.MethodGet, endpoint: "/api/v1/computers"},
			want: nil,
		},
		{
			name: "HEAD sends no bytes",
			args: args{body: nil, method: http.MethodHead, endpoint: "/api/v1/computers"},
			want: nil,
		},
		{
			name: "POST marshals a nil body as JSON",
			args: args{body: nil, method: http.MethodPost, endpoint: "/api/v1/computers"},
			want: []byte("null"),
		},
		{
			name: "POST marshals a JSON body for API paths",
			args: args{body: struct {
				Name string `json:"name"`
			}{Name: "example"}, method: http.MethodPost, endpoint: "/api/v1/computers"},
			want: []byte(`{"name":"example"}`),
		},
		{
			name: "POST marshals an XML body for JSSResource paths",
			args: args{body: xmlTestBody{Name: "example"}, method: http.MethodPost, endpoint: "/JSSResource/computers"},
			want: []byte(`<xmlTestBody><name>example</name></xmlTestBody>`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Integration{
				JamfProFQDN:          tt.fields.jamfProFQDN,
				AuthMethodDescriptor: tt.fields.AuthMethodDescriptor,
				Sugar:                tt.fields.Sugar,
				auth:                 tt.fields.auth,
				http:                 tt.fields.http,
			}
			got, err := j.marshalRequest(tt.args.body, tt.args.method, tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("Integration.marshalRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integration.marshalRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
