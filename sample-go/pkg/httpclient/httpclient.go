package httpclient

import (
	"context"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"strings"
)

const CorrelationIDHeaderName = "X-Correlationid"

// ServiceClient is an interface that can be used to issue http requests.
type ServiceClient interface {
	// NewRequest returns a request with a common set of configurations like:
	//  CorrelationID
	//  Technical token if client was created with NewOAuthClient()
	//  Retry settings
	NewRequest(context.Context) *resty.Request

	// GetHTTPClient returns the http client used by this ServiceClient.
	// It should be used for tests and mocks only.
	GetHTTPClient() *http.Client
}

type oauthClient struct {
	client *resty.Client
}

func (s *oauthClient) NewRequest(ctx context.Context) *resty.Request {
	request := s.client.R()
	request.SetContext(ctx)
	return request
}

func (s *oauthClient) GetHTTPClient() *http.Client {
	return s.client.GetClient()
}

// NewOAuthClient returns a new http client that can execute requests by authenticating against a given oauth endpoint
func NewOAuthClient(baseURL string, clientID string, clientSecret string, tokenURL string) ServiceClient {
	if !strings.HasSuffix(tokenURL, "/oauth/token") {
		tokenURL = strings.TrimSuffix(tokenURL, "/") + "/oauth/token"
	}
	creds := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}
	authClient := creds.Client(context.Background())

	client := resty.NewWithClient(authClient)
	client.SetBaseURL(baseURL)
	return &oauthClient{
		client: client,
	}
}
