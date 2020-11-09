package main

import (
	"fmt"
	"net/http"
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
	query := r.FormValue("q")
	isCheck := r.FormValue("check") == "on"
	if len(query) <= 0 {
		w.Write([]byte(html))
		return
	}

	w.Write([]byte(html + "<br/><h1>" + query + "</h1>" + fmt.Sprintf("%v", isCheck)))
	return
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8000", nil)
}
