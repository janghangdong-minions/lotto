package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagHTTPPort = flag.String("http", "", "webservice port number")
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
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
