package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("img")))) //currentl directory
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my first golang webpage`)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Welcome to my about page`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "<img src=\"mostain.jpg\" />")
}
