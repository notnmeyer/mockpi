package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerValidRequest(t *testing.T) {
	expectedContentType := "application/json; charset=utf-8"
	expectedResponseBody := `{"foo":"bar"}`
	expectedResponseCode := "418" // http.StatusTeapot

	// set up the request
	req, err := http.NewRequest("POST", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("X-Response-Json", expectedResponseBody)
	req.Header.Add("X-Response-Code", expectedResponseCode)

	// make the request
	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)
	h.ServeHTTP(rr, req)

	// status code
	if status := rr.Code; status != http.StatusTeapot {
		t.Errorf("handler returned wrong status code: got '%d' want '%d'", status, http.StatusOK)
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

// TODO: test invalid x-response-json
// TODO: test invalid x-response-code

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
