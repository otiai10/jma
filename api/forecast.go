package api

import (
	"time"
)

type (

	// 週間予報APIのレスポンス配列の要素. length = 2
	// それぞれの要素が何者であるか,
	// Keyの存在を見ないと判別できない.
	//
	//	第1要素 = ComprehensiveForecastEntry{
	//		TimeSeries[
	//			(a) 直近2日間/24時間単位の、天気・風・波,
	//			(b) 直近2日間/6時間単位の、降水確率,
	//			(c) 直近2日間の、朝の最低気温・日中の最高気温,
	//		]
	//	第2要素 = ComprehensiveForecastEntry{
	//		TimeSeries[
	//			(a) 直近1週間の、天気・降水確率・予報信頼性,
	//			(b) 直近1週間の、最高気温・最低気温,
	//		],
	// 		(c) 最高気温・最低気温の（おそらく1週間の）平均,
	//		(d) 最大降水量・最小降水量の（おそらく1週間の）平均,
	//	}

	ComprehensiveForecastEntry struct {
		// 第1要素・第2要素共通普遍のフィールド
		PublishingOffice string
		ReportDatetime   time.Time

		// 第1要素・第2要素共通のKeyだが,
		// 構造が異なるため, 包括的structにmapし,
		// 内容を見て何のデータか判断する必要がある.
		TimeSeries []ComprehensiveTimeSeries

		// 第2要素にのみ出現するフィールド
		TempAverage   TempAverage
		PrecipAverage PrecipAverage
	}

	// このエントリシリーズが何者であるかの識別は,
	// ComprehensiveAreaEntryの中身を見ないとわからない.
	ComprehensiveTimeSeries struct {
		// 当エントリシリーズの時間軸の定義.
		TimeDefines []time.Time

		// 地域ごとの,実際の値を格納している.
		// Areaの定義は「東京」「東京地方」など,
		// 必ずしも一意ではないことに注意.
		Areas []ComprehensiveAreaEntry
	}

	ComprehensiveAreaEntry struct {
		Area struct {
			Name string
			Code string
		}

		// 1-a)
		Weathers []string
		Winds    []string
		Waves    []string

		// 1-a / 2-a)
		WeatherCodes []string

		// 1-b / 2-a) 降水確率
		Pops []string // Probability of Precipitationの略だと思う.

		// 1-c) 朝の最低気温と日中の最高気温
		Temps []string

		// 2-a) 予報信頼性
		Reliabilities []string

		// 2-b) 最高気温・最低気温の予想
		TempsMin      []string
		TempsMinUpper []string
		TempsMinLower []string
		TempsMax      []string
		TempsMaxUpper []string
		TempsMaxLower []string
	}

	// 2-c) 地域平均気温（おそらく1週間）
	TempAverage struct {
		Areas []struct {
			Area struct {
				Name string
				Code string
			}
			Min string
			Max string
		}
	}

	// 2-d) 地域降水量（おそらく1週間）
	PrecipAverage struct {
		Areas []struct {
			Area struct {
				Name string
				Code string
			}
			Min string
			Max string
		}
	}
)
