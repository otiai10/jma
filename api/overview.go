package api

import "time"

type (
	OverviewWeek struct {
		PublishingOffice string
		ReportDatetime   time.Time
		HeadTitle        string
		Text             string
	}

	OverviewForecast struct {
		PublishingOffice string
		ReportDatetime   time.Time
		TargetArea       string
		HeadlineText     string
		Text             string
	}
)
