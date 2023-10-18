package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		responseBody := r.Header["X-Response-Json"][0]

		var responseCode int
		if val, ok := r.Header["X-Response-Code"]; !ok {
			responseCode = 200
		} else {
			var err error
			responseCode, err = strconv.Atoi(val[0])
			if err != nil {
				fmt.Println("not a value status code: ", err)
			}
		}

		w.WriteHeader(responseCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseBody))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error: ", err)
		os.Exit(1)
	}
}
