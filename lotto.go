package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
)

var (
	// TEMPLATES 는 lotto에서 사용하는 템플릿 글로벌 변수이다.
	TEMPLATES = template.New("")

	flagHTTPPort = flag.String("http", "", "webservice port number")
	flagGenNum   = flag.Bool("gennum", false, "generating random numbers")
	flagRangeMin = flag.Int("rmin", 1, "min value for generating random numbers")
	flagRangeMax = flag.Int("rmax", 45, "max value for generating random numbers")
	flagRange    = flag.Int("r", 7, "range for generating random numbers")
)

func main() {
	flag.Parse()

	if *flagHTTPPort != "" {
		ip, err := serviceIP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Service start: http://%s%s\n", ip, *flagHTTPPort)
		webserver()
	} else if *flagGenNum {
		if *flagRangeMin < 1 {
			fmt.Println("min값은 1이상 입력해주세요")
		}
		if *flagRangeMax < 45 {
			fmt.Println("max값은 45이상 입력해주세요")
		}
		if *flagRange < 7 {
			fmt.Println("range값을 7이상 입력해주세요")
		}
		GenRandomNums(*flagRangeMin, *flagRangeMax, *flagRange)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
