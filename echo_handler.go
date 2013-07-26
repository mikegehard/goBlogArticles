package blogarticles

import (
	"net/http"
)

import ()

type EchoHandler struct{}

func (e EchoHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	sayParam := r.FormValue("say")

	if sayParam == "Nothing" {
		rw.WriteHeader(404)
	} else {
		rw.Write([]byte(sayParam))
	}
}
