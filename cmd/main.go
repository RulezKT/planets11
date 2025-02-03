package main

import (
	"fmt"

	"github.com/RulezKT/planets11"
)

func main() {

	const DIR = "files"

	// мое время -682470731.47  [ 1978, 5, 17, 12, 47, 0 ]
	date_in_seconds := float64(-682470731)

	pl11 := planets11.Pl11{}
	pl11.Load(DIR)
	fmt.Println("pl11", pl11.Calc(date_in_seconds))

}
