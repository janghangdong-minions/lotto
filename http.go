package main

import (
	"fmt"
	"html/template"
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
	i, err := template.ParseFiles(
		"assets/html/header.html",
		"assets/html/bae.html",
		"assets/html/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = i.ExecuteTemplate(w, "bae", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleMan(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", "Man")
}

func handleSol(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", "sol")
}

func handleKok(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles(
		"assets/html/header.html",
		"assets/html/kok.html",
		"assets/html/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = i.ExecuteTemplate(w, "kok", nil)
	if err != nil {
		log.Fatal(err)
	}
}
