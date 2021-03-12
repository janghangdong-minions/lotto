package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagBae = flag.Bool("bae", false, "name of bae")
	flagSol = flag.Bool("sol", false, "name of solip")
	flagkok = flag.Bool("kok", false, "nick of kokbee")
	flagMan = flag.Bool("Man", false, "name of ManJe")
)

func main() {
	flag.Parse()

	if *flagBae {
		fmt.Println("bae seo young")
		os.Exit(0)
	} else if *flagSol {
		fmt.Println("lim sol Leaf")
	} else if *flagMan {
		fmt.Println("Kang Man JE")
	} else if *flagkok {
		fmt.Println("nickname kokbee")
		os.Exit(0)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
