package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	log.Println("Your Request method at bar : ", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, r *http.Request) {
	log.Println("Your request method at barred : ", r.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8000", nil)
}
