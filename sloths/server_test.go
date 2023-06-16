package sloths

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleSlothfulMessage(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/sloth", handleSlothfulMessage)

	svr := httptest.NewServer(router)
	defer svr.Close()

	c := NewClient(http.DefaultClient, svr.URL)
	m, err := c.GetSlothfulMessage()
	if err != nil {
		t.Fatalf("error in GetSlothfulMessage: %v", err)
	}
	if m.Message != "Stay slothful!" {
		t.Errorf(
			`message %s should contain string "Sloth"`,
			m.Message,
		)
	}
}
