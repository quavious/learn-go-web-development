package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Handler struct {
}

type Any interface{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Println(w, r)
	err := r.ParseForm()
	if checkError(err) {
		log.Fatalln(err)
	}
	tpl, err := template.ParseGlob("templates/*")

	if checkError(err) {
		return
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func checkError(e error) bool {
	if e != nil {
		log.Println(e.Error())
		return true
	}
	return false
}

func main() {
	var handler Handler
	http.ListenAndServe("localhost:8000", &handler)
	fmt.Println("Server On 8000")
}
