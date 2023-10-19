package main

import "testing"

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
