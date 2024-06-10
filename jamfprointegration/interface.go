// jamfpro_api_handler.go

package jamfprointegration

import (
	"net/http"

	"github.com/deploymenttheory/go-api-http-client/logger"
)

// JamfAPIHandler implements the APIHandler interface for the Jamf Pro API.
type Integration struct {
	BaseDomain           string
	InstanceName         string // InstanceName is the name of the Jamf instance.
	AuthMethodDescriptor string
	Logger               logger.Logger
	auth                 authInterface
}

// Info

func (j *Integration) Domain() string {
	return j.BaseDomain
}

func (j *Integration) GetAuthMethodDescriptor() string {
	return j.AuthMethodDescriptor
}

// Utilities

func (j *Integration) CheckRefreshToken() error {
	return j.checkRefreshToken()
}

func (j *Integration) PrepRequestParamsAndAuth(req *http.Request) error {
	err := j.prepRequest(req)
	return err
}

func (j *Integration) PrepRequestBody(body interface{}, method string, endpoint string) ([]byte, error) {
	return j.marshalRequest(body, method, endpoint)
}

func (j *Integration) MarshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) {
	return j.marshalMultipartRequest(fields, files)
}
