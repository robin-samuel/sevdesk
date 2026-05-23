package sevdesk

import "net/http"

// Option configures a Client. Pass options to [New].
type Option func(*Client)

// WithHTTPClient sets the underlying HTTP client. Useful for injecting
// custom timeouts, transports, or test doubles.
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.client = h }
}

// WithBaseURL overrides the API base URL. Defaults to https://my.sevdesk.de/api/v1.
// The internal instance lives at http://sevdesk.local/api/v1.
func WithBaseURL(url string) Option {
	return func(c *Client) { c.baseURL = url }
}

// WithUserAgent sets the User-Agent header sent on every request.
func WithUserAgent(ua string) Option {
	return func(c *Client) { c.userAgent = ua }
}
