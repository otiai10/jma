//go:generate go run ./tools/main.go -gen -target area -i https://www.jma.go.jp/bosai/common/const/area.json -o offices.go
package jma

import "strings"

type (
	Area struct {
		Code       string   `json:"code"`
		Name       string   `json:"name"`
		NameEn     string   `json:"enName"`
		Kana       string   `json:"kana"`
		OfficeName string   `json:"officeName"`
		Parent     string   `json:"parent"`
		Children   []string `json:"children"`

		// 以下、気象庁からのデータには存在しないフィールド
		NameEnLower string `json:"-"` // 観測所検索に使用
	}
)

func SearchOffice(name string) (res []Area) {
	name = strings.ToLower(name)
	for _, o := range Offices {
		if o.NameEnLower == name {
			return []Area{o}
		}
		if strings.Contains(o.NameEnLower, name) || strings.Contains(name, o.NameEnLower) {
			res = append(res, o)
		}
	}
	return res
}
