package main

import (
	"html/template"
	"log"
	"net/http"
)

func webserver() {
	http.HandleFunc("/", handleInit)
	http.HandleFunc("/bae", handleBae)
	http.HandleFunc("/man", handleMan)
	http.HandleFunc("/sol", handleSol)
	http.HandleFunc("/kok", handleKok)

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	err := http.ListenAndServe(*flagHTTPPort, nil) //nil의 정확한 용도를 이슈로 만들어봅시당.
	if err != nil {
		log.Fatal(err)
	}
}

func handleInit(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles(
		"assets/html/header.html",
		"assets/html/init.html",
		"assets/html/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = i.ExecuteTemplate(w, "init", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleBae(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles(
		"assets/html/header.html",
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
	i, err := template.ParseFiles(
		"assets/html/header.html",
		"assets/html/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = i.ExecuteTemplate(w, "man", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleSol(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles(
		"assets/html/header.html",
		"assets/html/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = i.ExecuteTemplate(w, "sol", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleKok(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles(
		"assets/html/header.html",
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
