package main

import (
	"html/template"
	"log"
	"net/http"
)

func webserver() {
	http.HandleFunc("/", handleInit)

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
