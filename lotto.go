package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var (
	flagHTTPPort = flag.String("http", "", "webservice port number")
	flagGenNum   = flag.Bool("gennum", false, "generating random numbers")
	flagRangeMin = flag.Int("rmin", 1, "generating random numbers for min value")
	flagRangeMax = flag.Int("rmax", 45, "generating random numbers for max value")
	flagRange    = flag.Int("r", 3, "generating random numbers for range")
)

// GenRandomNums 랜덤 값 출력
func GenRandomNums(min int, max int, ranges int) {
	// 로또 1~45까지 숫자 중 7개의 숫자를 뽑는다.
	for i := min; i < ranges+1; i++ {
		fmt.Println(rand.Intn(max-min) + min)
	}

}

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
		if *flagRangeMin <= 0 {
			fmt.Println("min값은 1이상 입력해주세요")
		}
		if *flagRangeMax < 45 {
			fmt.Println("max값은 45이상 입력해주세요")
		}
		if *flagRange < 1 {
			fmt.Println("range값을 1이상 입력해주세요")
		}
		GenRandomNums(*flagRangeMin, *flagRangeMax, *flagRange)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
