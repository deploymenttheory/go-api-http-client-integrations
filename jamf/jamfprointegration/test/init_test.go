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

// API integration has a settable expiry period and is set to 60s
// Account tokens do not have this. I think the expiry is an hour.

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
			name: "all vars set correctly",
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
		{
			name: "no client id",
			args: args{
				jamfProFQDN:       os.Getenv(ENV_KEY_JAMFPRO_FQDN),
				clientId:          "",
				clientSecret:      os.Getenv(ENV_KEY_CLIENT_SECRET),
				bufferPeriod:      10 * time.Minute,
				hideSensitiveData: true,
				executor:          &httpclient.ProdExecutor{Client: &http.Client{}},
				Sugar:             newSugaredDevelopmentLogger(),
			},
			wantErr: true,
		},
		{
			name: "no client secret",
			args: args{
				jamfProFQDN:       os.Getenv(ENV_KEY_JAMFPRO_FQDN),
				clientId:          os.Getenv(ENV_KEY_CLIENT_ID),
				clientSecret:      "",
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
			// Only testing for error as cannot predict pointers inside Integration output. Error is enough to deem success.
		})
	}
}

func TestBuildWithBasicAuth(t *testing.T) {
	type args struct {
		jamfProFQDN       string
		Sugar             *zap.SugaredLogger
		bufferPeriod      time.Duration
		username          string
		password          string
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
			name: "all vars set correctly",
			args: args{
				jamfProFQDN:       os.Getenv(ENV_KEY_JAMFPRO_FQDN),
				username:          os.Getenv(ENV_KEY_USERNAME),
				password:          os.Getenv(ENV_KEY_PASSWORD),
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
				username:          os.Getenv(ENV_KEY_USERNAME),
				password:          os.Getenv(ENV_KEY_PASSWORD),
				bufferPeriod:      100 * time.Minute,
				hideSensitiveData: true,
				executor:          &httpclient.ProdExecutor{Client: &http.Client{}},
				Sugar:             newSugaredDevelopmentLogger(),
			},
			wantErr: true,
		},
		{
			name: "no username",
			args: args{
				jamfProFQDN:       os.Getenv(ENV_KEY_JAMFPRO_FQDN),
				username:          "",
				password:          os.Getenv(ENV_KEY_PASSWORD),
				bufferPeriod:      100 * time.Minute,
				hideSensitiveData: true,
				executor:          &httpclient.ProdExecutor{Client: &http.Client{}},
				Sugar:             newSugaredDevelopmentLogger(),
			},
			wantErr: true,
		},
		{
			name: "no password",
			args: args{
				jamfProFQDN:       os.Getenv(ENV_KEY_JAMFPRO_FQDN),
				username:          os.Getenv(ENV_KEY_USERNAME),
				password:          "",
				bufferPeriod:      100 * time.Minute,
				hideSensitiveData: true,
				executor:          &httpclient.ProdExecutor{Client: &http.Client{}},
				Sugar:             newSugaredDevelopmentLogger(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := jamfprointegration.BuildWithBasicAuth(tt.args.jamfProFQDN, tt.args.Sugar, tt.args.bufferPeriod, tt.args.username, tt.args.password, tt.args.hideSensitiveData, tt.args.executor)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildWithBasicAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Only testing for error as cannot predict pointers inside Integration output. Error is enough to deem success.
		})
	}
}
