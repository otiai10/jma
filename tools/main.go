package main

import (
	"flag"
	"fmt"

	"github.com/otiai10/jma/tools/genweather"
)

var (
	gen       bool
	gentarget string
	geninput  string
	genoutput string
)

func init() {
	flag.BoolVar(&gen, "gen", false, "generate")
	flag.StringVar(&gentarget, "target", "weather", "gen target")
	flag.StringVar(&geninput, "i", "", "input file")
	flag.StringVar(&genoutput, "o", "", "output file")
}

func main() {

	flag.Parse()
	switch {
	case gen:
		if err := genweather.Generate(geninput, genoutput); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("No subucommand given")
	}
}
