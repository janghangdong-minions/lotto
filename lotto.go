package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flagBae := flag.Bool("bae", false, "name of bae")
	flag.Parse()

	if *flagBae {
		fmt.Println("bae seo young")
		os.Exit(0)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
