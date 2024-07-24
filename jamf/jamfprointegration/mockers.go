package jamfprointegration

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func test_newTestRequest() *http.Request {
	out, _ := http.NewRequest("GET", "https://yourserver.jamfcloud.com/JSSresource/endpoint", nil)
	return out
}

func test_newSugaredLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}

func test_newMockauth() *mockauth {
	out := mockauth{}
	return &out
}

func test_MockIntegration() *Integration {
	out := Integration{
		BaseDomain:           "",
		AuthMethodDescriptor: "test",
		Sugar:                test_newSugaredLogger(),
		auth:                 test_newMockauth(),
	}
	return &out
}

func test_getSampleJson() string {
	data := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}{
		AccessToken: "AccessToken",
		ExpiresIn:   500,
	}

	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
