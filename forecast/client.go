package forecast

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/otiai10/jma"
)

const (
	// https://www.jma.go.jp/bosai/forecast/data/forecast/130000.json
	ForecastBaseURL string = "https://www.jma.go.jp/bosai/forecast"
)

var (
	punctuation              = regexp.MustCompile("[、。]")
	header                   = "▶"
	spacer                   = " "
	maxCharCountForLineBreak = 100
)

type (
	ForecastEntity struct {
		PublishingOffice string
		ReportDatetime   time.Time
		TimeSeries       []SeriesEntry
		TempAverage      TempAverage
		PrecipAverage    PrecipAverage
	}
	TempAverage struct {
		Areas []AreaEntry
	}
	PrecipAverage struct {
		Areas []AreaEntry
	}
	SeriesEntry struct {
		TimeDefines []time.Time
		Areas       []AreaEntry
	}
	AreaEntry struct {
		Area          AreaIdentity
		WeatherCodes  []string
		Weathers      []string
		Winds         []string
		Waves         []string
		Pops          []string
		Reliabilities []string
		TempsMin      []string
		TempsMinUpper []string
		TempsMinLower []string
		TempsMax      []string
		TempsMaxUpper []string
		TempsMaxLower []string

		// Only appears in tempAverage and precipAverage
		Min string
		Max string
	}
	AreaIdentity struct {
		Name string
		Code string
	}

	Overview struct {
		PublishingOffice string
		ReportDatetime   time.Time
		HeadTitle        string
		Text             string
	}
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// TODO: Fix
func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    ForecastBaseURL,
	}
}

func (client *Client) Forecast(city jma.City) ([]ForecastEntity, error) {

	city = adjustCityCode(city)

	endpoint := fmt.Sprintf("%s/%s/%s.json", client.BaseURL, "data/forecast", city.Code())
	res, err := client.HTTPClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %s: %s", err.Error(), endpoint)
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response from the server: %s: %s", res.Status, endpoint)
	}
	entities := []ForecastEntity{}
	if err := json.NewDecoder(res.Body).Decode(&entities); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s: %s", err.Error(), endpoint)
	}
	return entities, nil
}

func (client *Client) OverviewToday(city jma.City) (*Overview, error) {

	city = adjustCityCode(city)

	endpoint := fmt.Sprintf("%s/%s/%s.json", client.BaseURL, "data/overview_forecast", city.Code())
	res, err := client.HTTPClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %s: %s", err.Error(), endpoint)
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response from the server: %s: %s", res.Status, endpoint)
	}
	overview := &Overview{}
	if err := json.NewDecoder(res.Body).Decode(overview); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s: %s", err.Error(), endpoint)
	}
	return overview, nil
}

func adjustCityCode(city jma.City) jma.City {
	switch city {
	case jma.Okinawa:
		// https://www.jma.go.jp/bosai/forecast/#area_type=offices&area_code=471000
		return city + 1000
	default:
		return city
	}
}
