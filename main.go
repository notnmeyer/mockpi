package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	flag "github.com/spf13/pflag"
)

type config struct {
	port int
}

var c config

func init() {
	flag.IntVarP(&c.port, "port", "p", 8080, "the listen port")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", handler)
	listenAddr := fmt.Sprintf(":%d", c.port)
	fmt.Printf("Listening on %s...\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		fmt.Println("server error: ", err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		contentType  = "application/json; charset=utf-8"
		responseBody = r.Header["X-Response-Json"][0]
		responseCode int
	)

	responseCode, err := validateResponseCode(r.Header)
	if err != nil {
		responseBody = err.Error()
	} else if !isJSON(responseBody) {
		responseBody = errorResponseFormatter("x-response-json must be valid JSON").Error()
		responseCode = http.StatusBadRequest
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(responseCode)
	w.Write([]byte(responseBody))
}

func validateResponseCode(header map[string][]string) (int, error) {
	if val, exists := header["X-Response-Code"]; exists {
		// verify its a number
		code, err := strconv.Atoi(val[0])
		if err != nil {
			return http.StatusBadRequest, errorResponseFormatter("x-response-code must be a number")
		}

		// verify it falls in the range of status codes
		if !(code >= 100 && code <= 599) {
			return http.StatusBadRequest, errorResponseFormatter("x-response-code must be between 100 and 599")
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

func errorResponseFormatter(msg string) error {
	return fmt.Errorf(`{"error":"%s"}`, msg)
}
