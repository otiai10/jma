package jma

import (
	"fmt"
	"strings"
)

// 都道府県コード
// http://www.kenchikushikai.or.jp/touroku/documents/code-todoufuken.pdf
// https://nlftp.mlit.go.jp/ksj/gml/codelist/PrefCd.html
type City int

func (cc City) Code() string {
	return fmt.Sprintf("%06d", cc)
}

const (
	Hokkaido  City = 10000
	Aomori    City = 20000
	Iwate     City = 30000
	Miyagi    City = 40000
	Akita     City = 50000
	Yamagata  City = 60000
	Fukushima City = 70000
	Ibaraki   City = 80000
	Tochigi   City = 90000
	Gunma     City = 100000
	Saitama   City = 110000
	Chiba     City = 120000
	Tokyo     City = 130000
	Kanagawa  City = 140000
	Niigata   City = 150000
	Toyama    City = 160000
	Ishikawa  City = 170000
	Fukui     City = 180000
	Yamanashi City = 190000
	Ngano     City = 200000
	Gifu      City = 210000
	Shizuoka  City = 220000
	Aichi     City = 230000
	Mie       City = 240000
	Shiga     City = 250000
	Kyoto     City = 260000
	Osaka     City = 270000
	Hyogo     City = 280000
	Nara      City = 290000
	Wakayama  City = 300000
	Tottori   City = 310000
	Shimane   City = 320000
	Okayama   City = 330000
	Hiroshima City = 340000
	Yamaguchi City = 350000
	Tokushima City = 360000
	Kagawa    City = 370000
	Ehime     City = 380000
	Kochi     City = 390000
	Fukuoka   City = 400000
	Saga      City = 410000
	Nagasaki  City = 420000
	Kumamoto  City = 430000
	Oita      City = 440000
	Miyazaki  City = 450000
	Kagoshima City = 460000
	Okinawa   City = 470000
)

func ParseCity(s string) (City, error) {
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
