package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flagBae := flag.Bool("bae", false, "name of bae")
	flagkok := flag.Bool("kok", false, "nick of kokbee")

	flag.Parse()

	if *flagBae {
		fmt.Println("bae seo young")
		os.Exit(0)
	} else if *flagkok {
		fmt.Println("nickname kokbee")
		os.Exit(0)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
