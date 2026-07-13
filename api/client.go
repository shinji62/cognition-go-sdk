package devin

import (
	"context"
	"fmt"
	"net/http"
)

// DefaultServer is the production base URL for the Devin External API.
const DefaultServer = "https://api.devin.ai"

// NewBearerClient returns a ClientWithResponses configured for the Devin
// production server and authenticated with the given API token. Every request
// is sent with an "Authorization: Bearer <token>" header.
//
// Pass additional ClientOption values (e.g. WithHTTPClient) to customize the
// underlying client.
func NewBearerClient(token string, opts ...ClientOption) (*ClientWithResponses, error) {
	return NewBearerClientWithServer(DefaultServer, token, opts...)
}

// NewBearerClientWithServer is like NewBearerClient but allows overriding the
// server base URL (useful for testing or non-production environments).
func NewBearerClientWithServer(server, token string, opts ...ClientOption) (*ClientWithResponses, error) {
	if token == "" {
		return nil, fmt.Errorf("devin: API token must not be empty")
	}
	authEditor := func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}
	opts = append([]ClientOption{WithRequestEditorFn(authEditor)}, opts...)
	return NewClientWithResponses(server, opts...)
}
