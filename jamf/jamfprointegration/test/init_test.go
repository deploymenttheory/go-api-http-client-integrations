package test

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/deploymenttheory/go-api-http-client-integrations/jamf/jamfprointegration"
	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"go.uber.org/zap"
)

const (
	ENV_KEY_JAMFPRO_FQDN  = "TEST_JAMFPRO_FQDN"
	ENV_KEY_CLIENT_ID     = "TEST_JAMFPRO_CLIENT_ID"
	ENV_KEY_CLIENT_SECRET = "TEST_JAMFPRO_CLIENT_SECRET"
	ENV_KEY_USERNAME      = "TEST_JAMFPRO_USERNAME"
	ENV_KEY_PASSWORD      = "TEST_JAMFPRO_PASSWORD"
)

func Test_BuildWithOAuth(t *testing.T) {
	type args struct {
		jamfProFQDN       string
		Sugar             *zap.SugaredLogger
		bufferPeriod      time.Duration
		clientId          string
		clientSecret      string
		hideSensitiveData bool
		executor          httpclient.HTTPExecutor
	}
	tests := []struct {
		name    string
		args    args
		want    *jamfprointegration.Integration
		wantErr bool
	}{
		{
			name: "All vars set correctly",
			args: args{
				jamfProFQDN:       os.Getenv(ENV_KEY_JAMFPRO_FQDN),
				clientId:          os.Getenv(ENV_KEY_CLIENT_ID),
				clientSecret:      os.Getenv(ENV_KEY_CLIENT_SECRET),
				bufferPeriod:      10 * time.Second,
				hideSensitiveData: true,
				executor:          &httpclient.ProdExecutor{Client: &http.Client{}},
				Sugar:             newSugaredDevelopmentLogger(),
			},
			wantErr: false,
		},
		{
			name: "buffer period too long",
			args: args{
				jamfProFQDN:       os.Getenv(ENV_KEY_JAMFPRO_FQDN),
				clientId:          os.Getenv(ENV_KEY_CLIENT_ID),
				clientSecret:      os.Getenv(ENV_KEY_CLIENT_SECRET),
				bufferPeriod:      10 * time.Minute,
				hideSensitiveData: true,
				executor:          &httpclient.ProdExecutor{Client: &http.Client{}},
				Sugar:             newSugaredDevelopmentLogger(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := jamfprointegration.BuildWithOAuth(tt.args.jamfProFQDN, tt.args.Sugar, tt.args.bufferPeriod, tt.args.clientId, tt.args.clientSecret, tt.args.hideSensitiveData, tt.args.executor)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildWithOAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("BuildWithOAuth() = %v, want %v", got, tt.want)
			// }
		})
	}
}
