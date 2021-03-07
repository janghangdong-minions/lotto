package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flagBae := flag.Bool("bae", false, "name of bae")
	flagSol := flag.Bool("sol", false, "name of solip")
	flag.Parse()

	if *flagBae {
		fmt.Println("bae seo young")
		os.Exit(0)
	} else if *flagSol {
		fmt.Println("lim sol Leaf")
		os.Exit(0)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
