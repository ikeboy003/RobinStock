package helpers

import (
	"RobinStock/globals"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func prepareURL(urlStr string, params map[string]string) (*url.URL, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	return u, nil
}

func makeRequest(urlStr string) (*http.Response, error) {
	resp, err := globals.Session.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	return resp, nil
}

func processJSONResponse(body io.ReadCloser, dataType string) (interface{}, error) {
	var data map[string]interface{}
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}
	return processData(data, dataType)
}

func processData(data map[string]interface{}, dataType string) (interface{}, error) {
	switch dataType {
	case "results", "pagination", "indexzero":
		results, ok := data["results"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("'results' not found in JSON response")
		}
		if dataType == "indexzero" && len(results) == 0 {
			return nil, fmt.Errorf("'results' is empty in JSON response")
		}
		if dataType == "pagination" {
			results, err := fetchNextPages(data, results)
			if err != nil {
				return nil, err
			}
			return results, nil
		}
		return results, nil
	default:
		return data, nil
	}
}

func fetchNextPages(data map[string]interface{}, results []interface{}) ([]interface{}, error) {
	nextPage, ok := data["next"].(string)
	for ok && nextPage != "" {
		resp, err := makeRequest(nextPage)
		if err != nil {
			return nil, fmt.Errorf("error fetching next page: %w", err)
		}
		defer resp.Body.Close()

		var pageData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&pageData); err != nil {
			return nil, fmt.Errorf("error decoding page JSON: %w", err)
		}

		nextPageResults, ok := pageData["results"].([]interface{})
		if !ok {
			break
		}
		results = append(results, nextPageResults...)
		nextPage, ok = pageData["next"].(string)

		if !ok && nextPage == "" {
			fmt.Println("No more pages to fetch")
		}
	}
	return results, nil
}

func readResponseBody(body io.ReadCloser) (interface{}, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	return bodyBytes, nil
}
