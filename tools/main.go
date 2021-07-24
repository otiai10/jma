package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/otiai10/jma/tools/gens"
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
		switch gentarget {
		case "weather":
			if err := gens.GenerateWeather(geninput, genoutput); err != nil {
				log.Fatalf("gen failed: %s %s: %v", gentarget, geninput, err)
			} else {
				fmt.Printf("DONE: %s %s\n", gentarget, genoutput)
			}
		case "area":
			if err := gens.GenerateArea(geninput, genoutput); err != nil {
				log.Fatalf("gen failed: %s %s: %v", gentarget, geninput, err)
			} else {
				fmt.Printf("DONE: %s %s\n", gentarget, genoutput)
			}
		default:
			fmt.Println("No generate target given")
		}
	default:
		fmt.Println("No subucommand given")
	}
}
