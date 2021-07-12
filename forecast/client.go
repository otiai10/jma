package forecast

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/otiai10/jma"
)

const (
	// https://www.jma.go.jp/bosai/forecast/data/forecast/130000.json
	ForecastBaseURL string = "https://www.jma.go.jp/bosai/forecast"
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

func (client *Client) ForecastByCity(city jma.CityCode) ([]ForecastEntity, error) {
	endpoint := fmt.Sprintf("%s/%s/%d.json", ForecastBaseURL, "data/forecast", city)
	res, err := client.HTTPClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	entities := []ForecastEntity{}
	if err := json.NewDecoder(res.Body).Decode(&entities); err != nil {
		return nil, err
	}
	return entities, nil
}
