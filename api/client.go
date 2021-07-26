package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DefaultForcastBaseURL = "https://www.jma.go.jp/bosai/forecast"
)

type Client struct {
	BaseURL string
}

func (c *Client) Forecast(areacode string) ([]ComprehensiveForecastEntry, error) {
	if c.BaseURL == "" {
		c.BaseURL = DefaultForcastBaseURL
	}
	endpoint := fmt.Sprintf("%s/%s/%s.json", c.BaseURL, "data/forecast", areacode)
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %s: %v", c.BaseURL, err)
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("got non-200 response: %s: %s", c.BaseURL, res.Status)
	}
	entries := []ComprehensiveForecastEntry{}
	if err := json.NewDecoder(res.Body).Decode(&entries); err != nil {
		return nil, fmt.Errorf("failed to decode response to struct: %s: %v", c.BaseURL, err)
	}
	return entries, nil
}
