package helpers

import (
	"RobinStock/globals"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// GetRequest makes a GET request to a given URL with optional parameters and returns the data.
func GetRequest(urlStr string, dataType string, params map[string]string, jsonifyData bool) (interface{}, error) {
	u, err := prepareURL(urlStr, params)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequest(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if jsonifyData {
		return processJSONResponse(resp.Body, dataType)
	}

	return readResponseBody(resp.Body)
}

// RequestPost makes a POST request to the specified URL with the given payload and returns the response.
func RequestPost(url string, payload map[string]interface{}, timeout int, setJSON bool, jsonifyData bool) (interface{}, error) {
	var body io.Reader
	var err error

	// Prepare the payload
	if setJSON {
		globals.Session.DefaultHeaders["Content-Type"] = "application/json"
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonPayload)
	} else {
		// Prepare form data if not JSON
		formData := make([]string, 0, len(payload))
		for key, value := range payload {
			formData = append(formData, fmt.Sprintf("%s=%v", key, value))
		}
		body = strings.NewReader(strings.Join(formData, "&"))
	}

	// Set the timeout for the request
	globals.Session.Client.Timeout = time.Duration(timeout) * time.Second

	// Make the request
	resp, err := globals.Session.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for JSON response and decode
	if jsonifyData {
		var responseData interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}

	// Return raw response if jsonifyData is false
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

// GetOutput returns the current global output stream.
func GetOutput() io.Writer {
	return globals.Output
}

// SetLoginState sets the login state.
func SetLoginState(loggedIn bool) {
	globals.LoggedIn = loggedIn
}

// SetOutput sets the global output stream.
func SetOutput(output io.Writer) {
	globals.Output = output
}

// LoginRequired checks if the user is logged in before calling the function.
func LoginRequired(f func() error) func() error {
	return func() error {
		if !globals.LoggedIn {
			return fmt.Errorf("must be logged in to call this function")
		}
		return f()
	}
}

// ConvertNilToString converts a nil return value to an empty string.
func ConvertNilToString(f func() (string, error)) func() (string, error) {
	return func() (string, error) {
		result, err := f()
		if err != nil {
			return "", err
		}
		if result == "" {
			return "", nil
		}
		return result, nil
	}
}

// FilterData extracts values from data based on the info key.
// data can be a slice of maps or a single map.
func FilterData(data interface{}, info string) interface{} {
	if data == nil {
		return data
	}

	switch d := data.(type) {
	case []map[string]interface{}: // Equivalent to Python's list of dicts
		if len(d) == 0 {
			return []interface{}{}
		}
		var result []interface{}
		for _, item := range d {
			if value, ok := item[info]; ok {
				result = append(result, value)
			} else {
				fmt.Fprintln(globals.Output, "Error: argument not key in dictionary")
				return []interface{}{}
			}
		}
		return result

	case map[string]interface{}: // Equivalent to Python's dict
		if value, ok := d[info]; ok {
			return value
		} else {
			fmt.Fprintln(globals.Output, "Error: argument not key in dictionary")
			return nil
		}

	default:
		fmt.Fprintln(globals.Output, "Error: unsupported data type")
		return nil
	}
}

// UpdateSession updates the session with the given headers.
func UpdateSession(key, value string) {
	globals.Session.DefaultHeaders[key] = value
}
