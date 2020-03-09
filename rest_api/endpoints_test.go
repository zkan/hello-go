package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/entries", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEntries)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":1,"first_name":"Kan","last_name":"Ouivirach","email_address":"test@example.com"}]`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}
