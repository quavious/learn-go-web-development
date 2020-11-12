package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var html string

func init() {
	html = `
		<html>
			<head>
				<meta charset="utf-8" />
				<title>Hello Go!</title>
			</head>
			<body>
				<form action="/" method="POST">
					<input name="q" type="text" id="text" />
					<input type="checkbox" id="check" name="check" />
					<button type="submit">Submit!</button>
				</form>
			</body>
		</html>
	`
}

func foo(w http.ResponseWriter, r *http.Request) {
	var q string
	fmt.Println(r.Method)
	if r.Method == "POST" {
		f, h, err := r.FormFile("q")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fmt.Println("\nfile: ", f, "\nheader: ", h, "\nerror", err)

		cp := bytes.NewBuffer(nil)
		n, err := io.Copy(cp, f)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		q = string(n)

		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(cp.Bytes())

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
		<form method="post", enctype="multipart/form-data">
			<input type="file" name="q" />
			<button type="submit">Submit</button>
		</form><br/>
	` + q))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8000", nil)
}
