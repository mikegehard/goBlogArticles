package blogarticles


import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
)

func TestEchosContent (t *testing.T) {
	handler := new(EchoHandler)
	expectedBody := "Hello"

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", fmt.Sprintf("http://example.com/echo?say=%s", expectedBody), nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	handler.ServeHTTP(recorder, req)

	switch recorder.Body.String() {
	case expectedBody:
		// body is equal so no need to do anything
	default:
		t.Errorf("Body (%s) did not match expectation (%s).", recorder.Body.String(), expectedBody)
	}
}

func TestReturns404IfYouSayNothing (t *testing.T) {
	handler := new(EchoHandler)

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://example.com/echo?say=Nothing", nil)
	if err != nil {
		t.Errorf("Failed to create request.")
	}

	handler.ServeHTTP(recorder, req)

	if recorder.Code != 404 {
		t.Errorf("Did not get a 404.")
	}
}

func TestUsingIntegrationStyle (t *testing.T) {
	server := httptest.NewServer(new(EchoHandler))
	defer server.Close()

	resp, err := http.DefaultClient.Get(fmt.Sprintf("%s?say=Nothing", server.URL))
	if err != nil {
		t.Errorf("Error performing request.")
	}

	if resp.StatusCode != 404 {
		t.Errorf("Did not get a 404.")
	}
}
