package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func getFlag() {
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
