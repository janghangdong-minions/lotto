package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/shurcooL/httpfs/html/vfstemplate"
)

// LoadTemplates 함수는 템플릿을 로딩합니다.
// httpfs라이브러리를 이용합니다.
// vfsgen으로 만들어진 VFS 폴더 내의 html파일과 빈 template을 연결합니다.
func LoadTemplates() (*template.Template, error) {
	t := template.New("").Funcs(funcMap)
	t, err := vfstemplate.ParseGlob(assets, t, "/html/*.html")
	// assets는 assets_generate.go의 결과파일인 assets_vfsdata.go안에 생기는 변수입니다. VFS 구조를 담고 있습니다.
	return t, err
}

var funcMap = template.FuncMap{
	"PaddingNum": PaddingNum,
}

func webserver() {
	// 템플릿 로딩을 위해서 vfs(가상파일시스템)을 로딩합니다.
	vfsTemplate, err := LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}
	TEMPLATES = vfsTemplate
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(assets)))
	http.HandleFunc("/", handleInit)
	http.HandleFunc("/rend", handleRendom)

	err = http.ListenAndServe(*flagHTTPPort, nil) //nil의 정확한 용도를 이슈로 만들어봅시당.
	if err != nil {
		log.Fatal(err)
	}
}
func handleInit(w http.ResponseWriter, r *http.Request) {

	l := GenRandomNums(*flagRangeMin, *flagRangeMax, *flagRange)

	w.Header().Set("Content-Type", "text/html")

	err := TEMPLATES.ExecuteTemplate(w, "init", l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
