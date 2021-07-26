package gens

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/otiai10/jma"
)

type (
	WeatherEntry    []string
	WeatherEntries  map[string]WeatherEntry
	SortableWeather struct {
		SortKey int
		Weather jma.Weather
	}
	SortableWeathers []SortableWeather
)

func doublequote(v interface{}) string {
	return fmt.Sprintf(`"%v"`, v)
}

func JSONEntryToWeather(code string, entry WeatherEntry) SortableWeather {
	i, _ := strconv.Atoi(code)
	return SortableWeather{
		SortKey: i,
		Weather: jma.Weather{
			Code:         code,
			IconSVGDay:   entry[0],
			IconSVGNight: entry[1],
			Category:     jma.Category(entry[2]),
			TextJa:       entry[3],
			TextEn:       entry[4],
		},
	}
}

func (s SortableWeathers) Len() int {
	return len(s)
}

func (s SortableWeathers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableWeathers) Less(i, j int) bool {
	return s[i].SortKey < s[j].SortKey
}

func WeatherToAstKeyValue(pos token.Pos, weather jma.Weather) *ast.KeyValueExpr {
	group := strings.Split(weather.IconSVGDay, ".")[0]
	emoji, ok := WeatherEmojiMap[group]
	if !ok {
		emoji.Slack = group
		emoji.Unicode = group
	}
	return &ast.KeyValueExpr{
		Key: &ast.BasicLit{
			Kind:     token.STRING,
			Value:    doublequote(weather.Code),
			ValuePos: pos,
		},
		Value: &ast.CompositeLit{
			Elts: []ast.Expr{
				&ast.KeyValueExpr{
					Key: ast.NewIdent("Code"),
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: doublequote(weather.Code),
					},
				},
				&ast.KeyValueExpr{
					Key: ast.NewIdent("Category"),
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: weather.Category.String(),
					},
				},
				&ast.KeyValueExpr{
					Key: ast.NewIdent("IconSVGDay"),
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: doublequote(weather.IconSVGDay),
					},
				},
				&ast.KeyValueExpr{
					Key: ast.NewIdent("IconSVGNight"),
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: doublequote(weather.IconSVGNight),
					},
				},
				&ast.KeyValueExpr{
					Key: ast.NewIdent("TextJa"),
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: doublequote(weather.TextJa),
					},
				},
				&ast.KeyValueExpr{
					Key: ast.NewIdent("TextEn"),
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: doublequote(weather.TextEn),
					},
				},
				&ast.KeyValueExpr{
					Key: ast.NewIdent("Emoji"),
					Value: &ast.CompositeLit{
						Type: ast.NewIdent("Emoji"),
						Elts: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: doublequote(emoji.Slack),
							},
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: doublequote(emoji.Unicode),
							},
						},
					},
				},
			},
		},
	}
}

func GenerateWeather(inputpath, outputpath string) error {

	input, err := os.Open(inputpath)
	if err != nil {
		return err
	}
	defer input.Close()
	entries := WeatherEntries{}
	if err := json.NewDecoder(input).Decode(&entries); err != nil {
		return err
	}

	output, err := os.Open(outputpath)
	if err != nil {
		return err
	}
	defer output.Close()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "weathers.go", output, parser.AllErrors)
	if err != nil {
		return err
	}

	sortables := SortableWeathers{}
	for code, entry := range entries {
		sortables = append(sortables, JSONEntryToWeather(code, entry))
	}

	sort.Sort(sortables)

	val := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0].(*ast.CompositeLit)
	// ast.Print(fset, val.Elts[0])
	line := fset.Position(val.Pos()).Line
	file := fset.File(1)
	val.Elts = []ast.Expr{}
	for _, w := range sortables {
		line += 1
		pos := file.LineStart(line)
		expr := WeatherToAstKeyValue(pos, w.Weather)
		val.Elts = append(val.Elts, expr)
	}

	// {{{ FIXME:
	buf := bytes.NewBuffer(nil)
	if err := format.Node(buf, fset, f); err != nil {
		return err
	}
	os.RemoveAll("./weathers.go")
	o, err := os.Create("./weathers.go")
	if err != nil {
		return err
	}
	if _, err := io.Copy(o, buf); err != nil {
		return err
	}
	// }}}

	return nil

}

var WeatherEmojiMap = map[string]struct {
	Slack   string
	Unicode string
}{
	"100": {":sunny:", "☀️"},
	"101": {":mostly_sunny:", "🌤"},
	"102": {":partly_sunny_rain:", "🌦"},
	"104": {":snow_cloud:", "🌨"},
	"110": {":partly_sunny:", "⛅️"},
	"112": {":partly_sunny_rain:", "🌦"},
	"115": {":snow_cloud:", "🌨"},
	"200": {":cloud:", "☁️"},
	"201": {":barely_sunny:", "🌥"},
	"202": {":rain_cloud:", "🌧"},
	"204": {":snow_cloud:", "🌨"},
	"210": {":partly_sunny:", "⛅️"},
	"212": {":rain_cloud:", "🌧"},
	"215": {":snow_cloud:", "🌨"},
	"300": {":umbrella:", "☂️"},
	"301": {":umbrella:", "☂️"},
	"302": {":umbrella:", "☂️"},
	"303": {":umbrella:", "☂️"},
	"308": {":umbrella_with_rain_drops:", "☔️"},
	"311": {":partly_sunny_rain:", "🌦"},
	"313": {":rain_cloud:", "🌧"},
	"314": {":snow_cloud:", "🌨"},
	"400": {":snowman_without_snow:", "⛄️"},
	"401": {":rain_cloud:", "🌧"},
	"402": {":snowflake:", "❄️"},
	"403": {":rain_cloud:", "🌨"},
	"406": {":snowman:", "☃️"},
	"411": {":snowflake:", "❄️"},
	"413": {":snowflake:", "❄️"},
	"414": {":rain_cloud:", "🌧"},
	// default:
	// 	return fmt.Sprintf(":grey_question: (%s)", group)
	// }
}
