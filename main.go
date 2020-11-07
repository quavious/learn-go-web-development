package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := "Index Page"
		tpl.ExecuteTemplate(w, "index.html", struct {
			Data string
		}{
			data,
		})
	}))

	http.Handle("/dog/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := "Dogs Page"
		tpl.ExecuteTemplate(w, "index.html", struct {
			Data string
		}{
			data,
		})
	}))

	http.Handle("/me/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idx := strings.Replace(r.URL.Path, "/me/", "", 1)
		data := fmt.Sprintf("My name is %s", idx)
		tpl.ExecuteTemplate(w, "index.html", struct {
			Data string
		}{
			data,
		})
	}))

	http.ListenAndServe("localhost:8000", nil)

}
