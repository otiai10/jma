package jma

import (
	"fmt"
	"strings"
)

// 都道府県コード
// http://www.kenchikushikai.or.jp/touroku/documents/code-todoufuken.pdf
// https://nlftp.mlit.go.jp/ksj/gml/codelist/PrefCd.html
type Prefecture int

func (pref Prefecture) Code() string {
	return fmt.Sprintf("%06d", pref)
}

const (
	Hokkaido  Prefecture = 10000
	Aomori    Prefecture = 20000
	Iwate     Prefecture = 30000
	Miyagi    Prefecture = 40000
	Akita     Prefecture = 50000
	Yamagata  Prefecture = 60000
	Fukushima Prefecture = 70000
	Ibaraki   Prefecture = 80000
	Tochigi   Prefecture = 90000
	Gunma     Prefecture = 100000
	Saitama   Prefecture = 110000
	Chiba     Prefecture = 120000
	Tokyo     Prefecture = 130000
	Kanagawa  Prefecture = 140000
	Niigata   Prefecture = 150000
	Toyama    Prefecture = 160000
	Ishikawa  Prefecture = 170000
	Fukui     Prefecture = 180000
	Yamanashi Prefecture = 190000
	Ngano     Prefecture = 200000
	Gifu      Prefecture = 210000
	Shizuoka  Prefecture = 220000
	Aichi     Prefecture = 230000
	Mie       Prefecture = 240000
	Shiga     Prefecture = 250000
	Kyoto     Prefecture = 260000
	Osaka     Prefecture = 270000
	Hyogo     Prefecture = 280000
	Nara      Prefecture = 290000
	Wakayama  Prefecture = 300000
	Tottori   Prefecture = 310000
	Shimane   Prefecture = 320000
	Okayama   Prefecture = 330000
	Hiroshima Prefecture = 340000
	Yamaguchi Prefecture = 350000
	Tokushima Prefecture = 360000
	Kagawa    Prefecture = 370000
	Ehime     Prefecture = 380000
	Kochi     Prefecture = 390000
	Fukuoka   Prefecture = 400000
	Saga      Prefecture = 410000
	Nagasaki  Prefecture = 420000
	Kumamoto  Prefecture = 430000
	Oita      Prefecture = 440000
	Miyazaki  Prefecture = 450000
	Kagoshima Prefecture = 460000
	Okinawa   Prefecture = 470000
)

func ParsePrefecture(s string) (Prefecture, error) {
	name := strings.Title(strings.ToLower(s))
	switch name {
	case "Hokkaido":
		return Hokkaido, nil
	case "Aomori":
		return Aomori, nil
	case "Iwate":
		return Iwate, nil
	case "Miyagi":
		return Miyagi, nil
	case "Akita":
		return Akita, nil
	case "Yamagata":
		return Yamagata, nil
	case "Fukushima":
		return Fukushima, nil
	case "Ibaraki":
		return Ibaraki, nil
	case "Tochigi":
		return Tochigi, nil
	case "Gunma":
		return Gunma, nil
	case "Saitama":
		return Saitama, nil
	case "Chiba":
		return Chiba, nil
	case "Tokyo":
		return Tokyo, nil
	case "Kanagawa":
		return Kanagawa, nil
	case "Niigata":
		return Niigata, nil
	case "Toyama":
		return Toyama, nil
	case "Ishikawa":
		return Ishikawa, nil
	case "Fukui":
		return Fukui, nil
	case "Yamanashi":
		return Yamanashi, nil
	case "Ngano":
		return Ngano, nil
	case "Gifu":
		return Gifu, nil
	case "Shizuoka":
		return Shizuoka, nil
	case "Aichi":
		return Aichi, nil
	case "Mie":
		return Mie, nil
	case "Shiga":
		return Shiga, nil
	case "Kyoto":
		return Kyoto, nil
	case "Osaka":
		return Osaka, nil
	case "Hyogo":
		return Hyogo, nil
	case "Nara":
		return Nara, nil
	case "Wakayama":
		return Wakayama, nil
	case "Tottori":
		return Tottori, nil
	case "Shimane":
		return Shimane, nil
	case "Okayama":
		return Okayama, nil
	case "Hiroshima":
		return Hiroshima, nil
	case "Yamaguchi":
		return Yamaguchi, nil
	case "Tokushima":
		return Tokushima, nil
	case "Kagawa":
		return Kagawa, nil
	case "Ehime":
		return Ehime, nil
	case "Kochi":
		return Kochi, nil
	case "Fukuoka":
		return Fukuoka, nil
	case "Saga":
		return Saga, nil
	case "Nagasaki":
		return Nagasaki, nil
	case "Kumamoto":
		return Kumamoto, nil
	case "Oita":
		return Oita, nil
	case "Miyazaki":
		return Miyazaki, nil
	case "Kagoshima":
		return Kagoshima, nil
	case "Okinawa":
		return Okinawa, nil
	default:
		return 0, fmt.Errorf("unknown city name: %s", name)
	}
}
