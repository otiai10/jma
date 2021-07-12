package main

import (
	"fmt"

	"github.com/otiai10/jma"
	"github.com/otiai10/jma/forecast"
)

func main() {
	client := forecast.NewClient()
	res, err := client.ForecastByCity(jma.Tokyo)
	fmt.Printf("%+v\n%v\n%v\n", res, len(res), err)
}
