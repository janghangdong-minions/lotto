package main

import (
	"flag"
)

var (
	flagHTTPPort = flag.String("http", "", "webservice port number")
	flagGenNum   = flag.Bool("gennum", false, "generating random numbers")
	flagRangeMin = flag.Int("rmin", 1, "min value for generating random numbers")
	flagRangeMax = flag.Int("rmax", 45, "max value for generating random numbers")
	flagRange    = flag.Int("r", 7, "range for generating random numbers")
)

func main() {
	getFlag()
}
