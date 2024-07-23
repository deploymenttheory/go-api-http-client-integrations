package jamfprointegration

import (
	"net/http"
	"testing"
)

func TestIntegration_prepRequest(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name        string
		integration *Integration
		args        args
		wantErr     bool
	}{
		{
			name:        "blank request",
			integration: test_MockIntegration(),
			args: args{
				req: test_newTestRequest(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.integration
			if err := j.prepRequest(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Integration.prepRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
