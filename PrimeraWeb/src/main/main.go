package main

import "net/http"

func main() {
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/martin", someFunc)
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hola Mundo"))
}
