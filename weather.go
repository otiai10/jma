//go:generate go run ./tools/main.go -gen -target weather -i ./data/weather.json -o weathers.go
package jma

import (
	"fmt"
	"strings"
)

type (
	Weather struct {
		Code         string
		Category     Category
		TextJa       string
		TextEn       string
		IconSVGDay   string // https://www.jma.go.jp/bosai/forecast/img/100.svg
		IconSVGNight string // https://www.jma.go.jp/bosai/forecast/img/500.svg
	}
	Category string
)

func (c Category) String() string {
	switch c {
	case Sunny:
		return "Sunny"
	case Cloudy:
		return "Cloudy"
	case Rainy:
		return "Rainy"
	case Snowy:
		return "Snowy"
	default:
		return "Unkonw"
	}
}

var (
	Sunny  Category = "100"
	Cloudy Category = "200"
	Rainy  Category = "300"
	Snowy  Category = "400"
)

// https://www.jma.go.jp/bosai/forecast/img/100.svg
func (w Weather) Icon(night ...bool) string {
	if len(night) != 0 && night[0] {
		return "https://www.jma.go.jp/bosai/forecast/img/" + w.IconSVGNight
	}
	return "https://www.jma.go.jp/bosai/forecast/img/" + w.IconSVGDay
}

func (w Weather) SlackEmoji() string {
	group := strings.Split(w.IconSVGDay, ".")[0]
	switch group {
	case "100":
		return ":sunny:"
	case "101":
		return ":mostly_sunny:"
	case "102":
		return ":partly_sunny_rain:"
	case "104":
		return ":snow_cloud:"
	case "110":
		return ":partly_sunny:"
	case "112":
		return ":partly_sunny_rain:"
	case "115":
		return ":snow_cloud:"
	case "200":
		return ":cloud:"
	case "201":
		return ":barely_sunny:"
	case "202":
		return ":rain_cloud:"
	case "204":
		return ":snow_cloud:"
	case "210":
		return ":partly_sunny:"
	case "212":
		return ":rain_cloud:"
	case "215":
		return ":snow_cloud:"
	case "300":
		return ":umbrella:"
	case "301":
		return ":umbrella:"
	case "302":
		return ":umbrella:"
	case "303":
		return ":umbrella:"
	case "308":
		return ":umbrella_with_rain_drops:"
	case "311":
		return ":partly_sunny_rain:"
	case "313":
		return ":rain_cloud:"
	case "314":
		return ":snow_cloud:"
	case "400":
		return ":snowman_without_snow:"
	case "401":
		return ":rain_cloud:"
	case "402":
		return ":snowflake:"
	case "403":
		return ":rain_cloud:"
	case "406":
		return ":snowman:"
	case "411":
		return ":snowflake:"
	case "413":
		return ":snowflake:"
	case "414":
		return ":rain_cloud:"
	default:
		return fmt.Sprintf(":grey_question: (%s)", group)
	}
}
