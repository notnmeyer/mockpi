package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerValidRequest(t *testing.T) {
	expectedContentType := "application/json; charset=utf-8"
	expectedResponseBody := `{"foo":"bar"}`
	expectedResponseCode := http.StatusTeapot

	// set up the request
	req, err := http.NewRequest("POST", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("X-Response-Json", expectedResponseBody)
	req.Header.Add("X-Response-Code", fmt.Sprint(expectedResponseCode))

	// make the request
	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)
	h.ServeHTTP(rr, req)

	// status code
	if status := rr.Code; status != expectedResponseCode {
		t.Errorf("handler returned wrong status code: got '%d' want '%d'", status, expectedResponseCode)
	}

	// response body
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got '%s' want '%s'", rr.Body.String(), expectedResponseBody)
	}

	// content-type
	if rr.Header()["Content-Type"][0] != expectedContentType {
		t.Errorf("handler returned wrong content-type: got '%s' wanted '%s'", rr.Header()["Content-Type"][0], expectedContentType)
	}
}

func TestHandlerWithInvalidXResponseJson(t *testing.T) {
	expectedContentType := "text/plain; charset=utf-8"
	invalidResponseBody := `{"foo":bar}`
	expectedResponseBody := "x-response-json must be valid JSON"
	expectedResponseCode := http.StatusBadRequest

	// set up the request
	req, err := http.NewRequest("POST", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("X-Response-Json", invalidResponseBody)
	req.Header.Add("X-Response-Code", fmt.Sprint(expectedResponseCode))

	// make the request
	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)
	h.ServeHTTP(rr, req)

	// status code
	if status := rr.Code; status != expectedResponseCode {
		t.Errorf("handler returned wrong status code: got '%d' want '%d'", status, expectedResponseCode)
	}

	// response body
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got '%s' want '%s'", rr.Body.String(), expectedResponseBody)
	}

	// content-type
	if rr.Header()["Content-Type"][0] != expectedContentType {
		t.Errorf("handler returned wrong content-type: got '%s' wanted '%s'", rr.Header()["Content-Type"][0], expectedContentType)
	}
}

func TestHandlerWithInvalidXResponseCode(t *testing.T) {
	expectedContentType := "text/plain; charset=utf-8"
	expectedResponseBody := "x-response-code must be a number\n"
	invalidResponseCode := "blah"
	expectedResponseCode := http.StatusBadRequest

	// set up the request
	req, err := http.NewRequest("POST", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("X-Response-Json", "{}")
	req.Header.Add("X-Response-Code", invalidResponseCode)

	// make the request
	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)
	h.ServeHTTP(rr, req)

	// status code
	if status := rr.Code; status != expectedResponseCode {
		t.Errorf("handler returned wrong status code: got '%d' want '%d'", status, expectedResponseCode)
	}

	// response body
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("handler returned unexpected body: got '%s' want '%s'", rr.Body.String(), expectedResponseBody)
	}

	// content-type
	if rr.Header()["Content-Type"][0] != expectedContentType {
		t.Errorf("handler returned wrong content-type: got '%s' wanted '%s'", rr.Header()["Content-Type"][0], expectedContentType)
	}
}

func TestValidateResponseCode(t *testing.T) {
	valid := map[string][]string{
		"X-Response-Code": {"201"},
	}
	if _, err := validateResponseCode(valid); err != nil {
		t.Errorf("expected %v to be a valid response code\n", valid)
	}

	invalid := map[string][]string{
		"X-Response-Code": {"0"},
	}
	if _, err := validateResponseCode(invalid); err == nil {
		t.Errorf("expected %v to be an invalid response code\n", invalid)
	}

	invalidStr := map[string][]string{
		"X-Response-Code": {"hello"},
	}
	if _, err := validateResponseCode(invalidStr); err == nil {
		t.Errorf("expected %v to be an invalid response code\n", invalid)
	}
}

func TestIsJSON(t *testing.T) {
	valid := `{"valid": "json"}`
	if !isJSON(valid) {
		t.Errorf("expected %s to be valid JSON\n", valid)
	}

	invalid := `{invalid: json,}`
	if isJSON(invalid) {
		t.Errorf("expected %s to be invalid JSON\n", invalid)
	}
}
