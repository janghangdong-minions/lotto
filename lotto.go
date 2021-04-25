package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	flagHTTPPort = flag.String("http", "", "webservice port number")
	flagGenNum   = flag.Bool("gennum", false, "generating random numbers")
	flagRangeMin = flag.Int("rmin", 1, "min value for generating random numbers")
	flagRangeMax = flag.Int("rmax", 45, "max value for generating random numbers")
	flagRange    = flag.Int("r", 7, "range for generating random numbers")
)

// GenRandomNums 랜덤 값 출력
func GenRandomNums(min, max, ranges int) LottoNum {
	// 로또 1~45까지 숫자 중 7개의 숫자를 뽑는다.
	var nums []int
	for i := 0; i < ranges; i++ {
		rand.Seed(time.Now().UnixNano())
		nums = append(nums, rand.Intn(max-min)+min)
	}

	l := new(LottoNum)
	l.BasicNums = nums

	return *l
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
