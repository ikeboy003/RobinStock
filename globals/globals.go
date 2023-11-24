package globals

import (
	"io"
	"net/http"
	"os"
)

var (
	LoggedIn bool
	Output   io.Writer = os.Stdout
	Session  *CustomClient
)

// CustomClient wraps http.Client to provide custom functionality.
type CustomClient struct {
	*http.Client
	DefaultHeaders map[string]string
}

// NewRequest creates and sends an HTTP request with default headers.
func (c *CustomClient) NewRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	// Set default headers for the request
	for key, value := range c.DefaultHeaders {
		req.Header.Set(key, value)
	}
	return c.Do(req)
}

func init() {
	LoggedIn = false
	Session = &CustomClient{
		Client: &http.Client{},
		DefaultHeaders: map[string]string{
			"Accept":                  "*/*",
			"Accept-Encoding":         "gzip,deflate,br",
			"Accept-Language":         "en-US,en;q=1",
			"Content-Type":            "application/x-www-form-urlencoded; charset=utf-8",
			"X-Robinhood-API-Version": "1.315.0",
			"Connection":              "keep-alive",
			"User-Agent":              "*",
		},
	}
}
