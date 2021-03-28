package main

import (
	"flag"
	"fmt"
	"log"
	_ "os"
)

var (
	flagHttpPort = flag.String("http", "", "Web Port Number")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("mim : ")
	flag.Parse()

	ip, err := serverData()
	if err != nil {
		log.Fatal(err)
	}
	if *flagHttpPort != "" {
		if *flagHttpPort == ":80" {
			fmt.Printf("Http://%s", ip)

		} else if *flagHttpPort == ":443" {
			fmt.Printf("Https://%s", ip)
		} else {
			fmt.Printf("Dev Mode\n")
			fmt.Printf("Http://%s%s\n", ip, *flagHttpPort)
		}
		webPage(*flagHttpPort)
	}
}
