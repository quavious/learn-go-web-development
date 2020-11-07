package main

import "net/http"

type h1 int
type h2 int

func (h *h1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dog Says Bark!"))
}

func (h *h2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cat Says Meow!"))
}

func r(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It is just a function, not handler."))
}

func main() {
	var p h1
	var q h2

	//mux := http.NewServeMux()
	http.Handle("/dog/", &p)
	http.Handle("/cat/", &q)

	http.HandleFunc("/func", r)

	http.ListenAndServe("localhost:8000", nil)
}
