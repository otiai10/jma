package jma

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestWeather_SlackEmoji(t *testing.T) {
	code := "101"
	w, ok := Weathers[code]
	Expect(t, ok).ToBe(true)
	Expect(t, w).TypeOf("jma.Weather")
	// Expect(t, w.SlackEmoji()).ToBe(":mostly_sunny:")

	// w = Weather{Code: "9999"}
	// Expect(t, w.SlackEmoji()).Match("question")

	// for _, w := range Weathers {
	// 	Expect(t, w.SlackEmoji()).Not().Match("question")
	// }
}
