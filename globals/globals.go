package globals

import (
	"net/http"
	"os"
)

var (
	LoggedIn       bool
	Output         = os.Stdout
	Session        *http.Client
	DefaultHeaders map[string]string
)

func init() {
	LoggedIn = false
	Output = os.Stdout
	Session = &http.Client{}
	DefaultHeaders = map[string]string{
		"Accept":                  "*/*",
		"Accept-Encoding":         "gzip,deflate,br",
		"Accept-Language":         "en-US,en;q=1",
		"Content-Type":            "application/x-www-form-urlencoded; charset=utf-8",
		"X-Robinhood-API-Version": "1.315.0",
		"Connection":              "keep-alive",
		"User-Agent":              "*",
	}
}
