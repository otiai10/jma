package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/otiai10/ja"
	"github.com/otiai10/jma"
	"github.com/otiai10/jma/forecast"
)

var (
	overview bool = false
	weekly   bool = false
)

func init() {
	flag.BoolVar(&overview, "overview", false, "Show overview")
	flag.BoolVar(&overview, "O", false, "Show overview (alias)")
	flag.BoolVar(&weekly, "week", false, "Show forecast of the week")
	flag.BoolVar(&weekly, "w", false, "Show forecast of the week (alias)")
}

func main() {
	flag.Parse()
	a := flag.Arg(0)
	if a == "" {
		a = "tokyo"
	}
	city, err := jma.ParsePrefecture(a)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if overview {
		if err := showOvereview(city); err != nil {
			log.Fatalln(err)
		}
		return
	}

	if err := showForecast(city); err != nil {
		log.Fatalln(err)
	}
}

func showForecast(city jma.Prefecture) error {
	client := forecast.NewClient()
	res, err := client.Forecast(city)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v\n", res[0])
	fmt.Printf("%+v\n", res[1])
	return nil
}

func showOvereview(city jma.Prefecture) error {
	client := forecast.NewClient()
	ov, err := client.OverviewToday(city)
	if err != nil {
		return err
	}
	if ov.HeadTitle != "" {
		fmt.Println(ov.HeadTitle)
	}
	for _, b := range ja.Paragraph(ov.Text, ja.Within(40)) {
		for i, line := range b {
			if i == 0 {
				fmt.Print("▶ ")
			} else {
				fmt.Print("  ")
			}
			fmt.Println(line)
		}
	}
	return nil
}
