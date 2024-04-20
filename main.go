package main

import "net/http"

func main() {
	http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))
}
