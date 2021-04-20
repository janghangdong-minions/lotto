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
	flagGenNum   = flag.Bool("gennum", false, "generating non-overlaped random number")
)

func GenRandomNums() {
	// 로또 1~45까지 숫자중 중복없이 7개의 숫자를 뽑는다.
	min := 1
	max := 45
	for i := 0; i < 7; i++ {
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
		GenRandomNums()
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
