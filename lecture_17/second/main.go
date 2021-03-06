package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", img)
	http.HandleFunc("/mostain.jpg", myPic)
	http.HandleFunc("/logemail", email)
	http.ListenAndServe(":8080", nil)
}

func img(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="mostain.jpg">	`)
}

func myPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "mostain.jpg")
}

func email(w http.ResponseWriter, req *http.Request) {

	fmt.Println(req.RemoteAddr)

	for key, val := range req.Header {
		fmt.Println(key, val)
	}
	http.ServeFile(w, req, "email_open_log_pic.gif")
}
