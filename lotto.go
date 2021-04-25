package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagHTTPPort = flag.String("http", "", "webservice port number")
	flagGenNum   = flag.Bool("gennum", false, "generating random numbers")
	flagRangeMin = flag.Int("rmin", 1, "min value for generating random numbers")
	flagRangeMax = flag.Int("rmax", 45, "max value for generating random numbers")
	flagRange    = flag.Int("r", 7, "range for generating random numbers")
)

func main() {
	flag.Parse()

	if *flagHTTPPort != "" {
		if *flagRange < 1 {
			fmt.Println("range값을 1이상 입력해주세요")
		}
		if *flagRangeMax <= *flagRangeMin {
			fmt.Println("rmax값은 rmin값보다 큰 값을 입력해주세요")
		}
		ip, err := serviceIP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Service start: http://%s%s\n", ip, *flagHTTPPort)
		webserver()
	} else if *flagGenNum {
		if *flagRange < 1 {
			fmt.Println("range값을 1이상 입력해주세요")
		}
		if *flagRangeMax <= *flagRangeMin {
			fmt.Println("rmax값은 rmin값보다 큰 값을 입력해주세요")
		}
		GenRandomNums(*flagRangeMin, *flagRangeMax, *flagRange)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
