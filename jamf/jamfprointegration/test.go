package jamfprointegration

import (
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

func test_newTestRequest() *http.Request {
	url, _ := url.Parse("https://yourserver.jamfcloud.com/JSSresource/endpoint")
	out := &http.Request{
		URL: url,
	}
	return out
}

func test_newSugaredLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
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
