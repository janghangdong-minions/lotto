package main

import (
	"html/template"
	"log"
	"net/http"
)

var funcMap = template.FuncMap{
	"PaddingNum": PaddingNum,
}

func webserver() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(assets)))
	http.HandleFunc("/", handleInit)

	err := http.ListenAndServe(*flagHTTPPort, nil) //nil의 정확한 용도를 이슈로 만들어봅시당.
	if err != nil {
		log.Fatal(err)
	}
}

func handleInit(w http.ResponseWriter, r *http.Request) {
	newTpl := template.New("").Funcs(funcMap)
	t, err := newTpl.ParseFiles(
		"assets/html/header.html",
		"assets/html/init.html",
		"assets/html/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	l := GenRandomNums(*flagRangeMin, *flagRangeMax, *flagRange)

	err = t.ExecuteTemplate(w, "init", l)
	if err != nil {
		log.Fatal(err)
	}
}
