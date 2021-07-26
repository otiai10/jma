package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/otiai10/jma"
	"github.com/otiai10/jma/api"
)

var (
	baseurl string
	list    bool
)

func init() {
	// flag.StringVar(&city, "city", "tokyo", "都市や地域の名前")
	flag.StringVar(&baseurl, "base", "https://www.jma.go.jp/bosai/forecast", "データソースURL")
	flag.BoolVar(&list, "list", false, "予報対象地域の一覧")
}

func main() {
	flag.Parse()

	if list {
		for _, o := range jma.Offices {
			fmt.Printf("%v:\t\t%v\n", o.NameEnLower, o.OfficeName)
		}
		return
	}

	city := flag.Arg(0)
	if city == "" {
		city = "tokyo"
	}
	areas := jma.SearchOffice(city)

	if len(areas) == 0 {
		fmt.Printf("名前「%s」に一致する地域がありませんでした. -list で地域一覧を確認してください.\n", city)
		os.Exit(1)
	}
	if len(areas) > 1 {
		fmt.Printf("%d個の候補地域が発見されました. いずれかのひとつを選んでください.\n", len(areas))
		for _, a := range areas {
			fmt.Printf("%s: %s (%s)\n", a.NameEnLower, a.Name, a.OfficeName)
		}
		return
	}

	client := &api.Client{BaseURL: baseurl}
	res, err := client.Forecast(areas[0].Code)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tpl, err := template.New("weekly_forecast").Funcs(template.FuncMap{
		"tfmt": func(t time.Time, format string) string { return t.Format(format) },
	}).Parse(weeklyForecastTemplate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// FIXME: resの構造が複雑であることに起因する整形をここでやるべきか？
	table := jma.ConvertWeeklyForecast(res[1].TimeSeries[0])
	tpl.Execute(os.Stdout, table)
}

var (
	weeklyForecastTemplate = `		{{range .TimeDefines}}{{tfmt . "01/02"}}	{{end}}
{{range .Areas}}【{{.Area.Name}}】	{{range .Entries}}{{.Weather.Emoji.Unicode}}  {{.PoP}}	{{end}}
{{end}}`
)
