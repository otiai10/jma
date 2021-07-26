//go:generate go run ./tools/main.go -gen -target weather -i ./data/weather.json -o weathers.go
package jma

type (
	Weather struct {
		Code         string
		Category     Category
		TextJa       string
		TextEn       string
		IconSVGDay   string // https://www.jma.go.jp/bosai/forecast/img/100.svg
		IconSVGNight string // https://www.jma.go.jp/bosai/forecast/img/500.svg
		Emoji        Emoji
	}
	Emoji struct {
		Slack   string
		Unicode string
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
