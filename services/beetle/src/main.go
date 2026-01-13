package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("beetle starting on :8080")

	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", nil)
}

