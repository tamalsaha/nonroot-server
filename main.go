package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("listening on :80")
	err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World"))
	}))
	if err != nil {
		panic(err)
	}
	fmt.Println("DONE!")
}
