package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", writeCookie)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/more", moreCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8000", nil)
}

func writeCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "MyCookie",
		Value: "Chocolate Cookies",
		Path:  "/",
	})
	fmt.Fprintf(w, "Cookie written\n")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("MyCookie")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, "Your Cookie 1 is : ", c1)
	}

	c2, err := r.Cookie("general")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, "Your Cookie 1 is : ", c2)
	}

	c3, err := r.Cookie("specific")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, "Your Cookie 1 is : ", c3)
	}
}

func moreCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "Some Other Cookie about general",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "Some Other Cookie about specific",
	})
	fmt.Fprintln(w, "More cookies written")
}
