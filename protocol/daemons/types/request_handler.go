package types

import (
	"context"
	"net/http"
	"os"
	"strings"
)

// RequestHandlerImpl is the struct that implements the `RequestHandler` interface.
type RequestHandlerImpl struct {
	client *http.Client
}

// RequestHandler is an interface that handles making HTTP requests.
type RequestHandler interface {
	Get(ctx context.Context, url string) (*http.Response, error)
}

// NewRequestHandlerImpl creates a new RequestHandlerImpl. It manages making HTTP requests.
func NewRequestHandlerImpl(client *http.Client) *RequestHandlerImpl {
	return &RequestHandlerImpl{
		client: client,
	}
}

// Get wraps `http.Get` which makes an HTTP GET request to a URL and returns a response.
func (r *RequestHandlerImpl) Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if strings.Contains(url, "apidojo-yahoo-finance-v1.p.rapidapi.com") {
		req.Header.Add("x-rapidapi-host", "apidojo-yahoo-finance-v1.p.rapidapi.com")
		apiKey := os.Getenv("YAHOO_FINANCE_API_KEY")
		if apiKey == "" {
			panic("YAHOO_FINANCE_API_KEY environment variable is not set")
		}
		req.Header.Add("x-rapidapi-key", apiKey)
	}

	return r.client.Do(req)
}
