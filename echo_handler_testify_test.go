package blogarticles


import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
)

func TestEchosContentTestify (t *testing.T) {
	handler := new(EchoHandler)
	expectedBody := "Hello"

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", fmt.Sprintf("http://example.com/echo?say=%s", expectedBody), nil)
	assert.Nil(t, err)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, expectedBody, recorder.Body.String())
}

func TestReturns404IfYouSayNothingTestify (t *testing.T) {
	handler := new(EchoHandler)

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://example.com/echo?say=Nothing", nil)
	assert.Nil(t, err)

	handler.ServeHTTP(recorder, req)

	assert.Equal(t, 404, recorder.Code)
}

func TestClientTestify (t *testing.T) {
	server := httptest.NewServer(new(EchoHandler))
	defer server.Close()

	// Pretend this is some sort of Go client...
	url := fmt.Sprintf("%s?say=Nothing", server.URL)
	resp, err := http.DefaultClient.Get(url)
	assert.Nil(t, err)

	assert.Equal(t, 404, resp.StatusCode)
}
