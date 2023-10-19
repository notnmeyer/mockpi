package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			contentType  = "application/json; charset=utf-8"
			responseBody = r.Header["X-Response-Json"][0]
			responseCode int
		)

		responseCode, err := validateResponseCode(r.Header)
		if err != nil {
			responseBody = err.Error()
			contentType = "text/plain; charset=utf-8"
		}

		if !isJSON(responseBody) {
			responseBody = fmt.Errorf("x-response-json must be valid JSON").Error()
			responseCode = http.StatusBadRequest
			contentType = "text/plain; charset=utf-8"
		}

		w.WriteHeader(responseCode)
		w.Header().Set("Content-Type", contentType)
		w.Write([]byte(responseBody))
	})

	listenAddr := ":8080"
	fmt.Printf("Listening on %s...\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		fmt.Println("server error: ", err)
		os.Exit(1)
	}
}

func validateResponseCode(header map[string][]string) (int, error) {
	if val, exists := header["X-Response-Code"]; exists {
		// verify its a number
		code, err := strconv.Atoi(val[0])
		if err != nil {
			return http.StatusBadRequest, fmt.Errorf("x-response-code must be a number\n")
		}

		// verify it falls in the range of status codes
		if !(code >= 100 && code <= 599) {
			return http.StatusBadRequest, fmt.Errorf("x-response-code must be between 100 and 599\n")
		}

		return code, nil
	}

	// if x-response-code was not supplied default to a 200
	return http.StatusOK, nil
}

func isJSON(s string) bool {
	var j json.RawMessage
	return json.Unmarshal([]byte(s), &j) == nil
}
