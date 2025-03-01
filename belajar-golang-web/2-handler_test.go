package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	/*
	 * @param {http.ResponseWriter} writer - response yang akan diberikan ke client
	 * @param {*http.Request} request - request yang akan diterima/dikirim dari client
	 */
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			panic(err)
		}

	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
