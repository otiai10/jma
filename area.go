//go:generate go run ./tools/main.go -gen -target area -i https://www.jma.go.jp/bosai/common/const/area.json -o offices.go
package jma

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
