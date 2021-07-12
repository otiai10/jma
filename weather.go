package jma

type (
	Weather struct {
		Code         int
		CodeText     string
		TextJa       string
		TextEn       string
		IconSVGDay   string // https://www.jma.go.jp/bosai/forecast/img/100.svg
		IconSVGNight string // https://www.jma.go.jp/bosai/forecast/img/500.svg
	}
)

var (
	WeatherDictionary = map[string]Weather{
		"100": {
			Code:         100,
			CodeText:     "100",
			TextJa:       "晴",
			TextEn:       "CLEAR",
			IconSVGDay:   "100.svg",
			IconSVGNight: "500.svg",
		},
		"101": {
			Code:         100,
			CodeText:     "100",
			IconSVGDay:   "101.svg",
			IconSVGNight: "501.svg",
			TextJa:       "晴時々曇",
			TextEn:       "PARTLY CLOUDY",
		},
	}
)
