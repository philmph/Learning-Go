package main

import (
	"flag"
	"slices"
	"strings"
)

func main() {
	activeArgs := flag.String("active", "", "Comma separated list of active snippets")
	flag.Parse()

	active := strings.Split(*activeArgs, ",")

	if slices.Contains(active, "slices") {
		slicesFunc()
	}

	if slices.Contains(active, "maps") {
		mapsFunc()
	}

	if slices.Contains(active, "loops") {
		loopsFunc()
	}
}
