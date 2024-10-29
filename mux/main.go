package main

import (
	"io"
	"net/http"
)

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/test", test)
	http.ListenAndServe(":5050", mux)
}

func test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string ("Ini halaman test"))
}