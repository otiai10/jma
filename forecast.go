package jma

import (
	"time"

	"github.com/otiai10/jma/api"
)

type (
	Reliability string
	PoP         string
)

type (
	// 直近1週間の天気予報のエントリ
	ForecastEntry struct {
		Time        time.Time
		Weather     Weather
		Reliability Reliability
		PoP         PoP
	}
	AreaForecast struct {
		Area    Area
		Entries []ForecastEntry
	}
	ForecastTable struct {
		TimeDefines []time.Time
		Areas       []AreaForecast
	}
)

func ConvertWeeklyForecast(series api.ComprehensiveTimeSeries) (table ForecastTable) {
	table.TimeDefines = series.TimeDefines
	for _, a := range series.Areas {
		af := AreaForecast{Area: Area{Code: a.Area.Code, Name: a.Area.Name}}
		for i, t := range series.TimeDefines {
			af.Entries = append(af.Entries, ForecastEntry{
				Time:        t,
				Weather:     Weathers[a.WeatherCodes[i]],
				Reliability: Reliability(a.Reliabilities[i]),
				PoP:         PoP(a.Pops[i]),
			})
		}
		table.Areas = append(table.Areas, af)
	}
	return
}
