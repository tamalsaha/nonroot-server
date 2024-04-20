package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("listening on :80")
	http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))
}
