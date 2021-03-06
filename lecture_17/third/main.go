package main

import (
	"io"
	"net/http"
)

func main() {

	//http.Handle("/", http.FileServer(http.Dir("assets")))
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("assets"))))

	//http.HandleFunc("/img", img)
	http.ListenAndServe(":8080", nil)
}

func img(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="img/assets/mostain.jpg">`)
}
