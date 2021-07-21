package jma

import "fmt"

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
