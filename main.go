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
	responseBody, responseCode := buildResponse(r.Header)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(responseCode)
	w.Write([]byte(responseBody))
}

func buildResponse(header map[string][]string) (string, int) {
	responseBody, err := validateResponseBody(header)
	if err != nil {
		return err.Error(), http.StatusBadRequest

	}

	responseCode, err := validateResponseCode(header)
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}

	return responseBody, responseCode
}

func validateResponseBody(header map[string][]string) (string, error) {
	if _, exists := header["X-Response-Json"]; !exists {
		return "", errorResponseFormatter("x-response-json must be set on the request")
	}

	if !isJSON(header["X-Response-Json"][0]) {
		return "", errorResponseFormatter("x-response-json must be valid JSON")
	}

	return header["X-Response-Json"][0], nil
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
