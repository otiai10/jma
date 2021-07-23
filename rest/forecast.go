package rest

import (
	"time"

	"github.com/otiai10/jma"
)

type (
	PoP         int
	Reliability string
	Weather     struct {
		Time time.Time
		jma.Weather
		Text        string
		Wind        string
		Wave        string
		PoP         PoP
		Reliability Reliability
	}
	Area struct {
		Name string // "東京地方"
		Code string // "130010"
	}
)

func ConvertForecastResult() {

}
