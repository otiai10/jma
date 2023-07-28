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

func (c *Client) OverviewWeek(areacode string) (*OverviewWeek, error) {
	return overview[OverviewWeek](c, "data/overview_week", areacode)
}

func (c *Client) Overview(areacode string) (*OverviewForecast, error) {
	return overview[OverviewForecast](c, "data/overview_forecast", areacode)
}

func overview[T any](c *Client, path, areacode string) (*T, error) {
	if c.BaseURL == "" {
		c.BaseURL = DefaultForcastBaseURL
	}
	endpoint := fmt.Sprintf("%s/%s/%s.json", c.BaseURL, path, areacode)
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ov T
	if err := json.NewDecoder(res.Body).Decode(&ov); err != nil {
		return nil, err
	}
	return &ov, nil

}
