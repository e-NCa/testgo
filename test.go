package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Count int
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := Page{"hello", 1}
	tmpl, err := template.ParseFiles("test.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
