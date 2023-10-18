package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// method := r.Method
		// path := r.URL.Path
		responseHeader := r.Header["X-Response-Json"][0]
		response := fmt.Sprint(responseHeader)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error: ", err)
		os.Exit(1)
	}
}
