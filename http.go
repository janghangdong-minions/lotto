package main

import (
	"fmt"
	"log"
	"net/http"
)

func webserver() {
	http.HandleFunc("/", handleInit)
	http.HandleFunc("/bae", handleBae)
	http.HandleFunc("/Man", handleMan)
	http.HandleFunc("/sol", handleSol)
	http.HandleFunc("/kok", handleKok)

	err := http.ListenAndServe(*flagHTTPPort, nil) //nil의 정확한 용도를 이슈로 만들어봅시당.
	if err != nil {
		log.Fatal(err)
	}
}

func handleInit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func handleBae(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", "bae")
}

func handleMan(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", "Man")
}

func handleSol(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", "sol")
}

func handleKok(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", "kok")
}
