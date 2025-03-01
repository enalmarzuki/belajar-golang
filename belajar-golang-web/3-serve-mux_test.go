package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, response *http.Request) {
		fmt.Fprint(writer, "Hello World!")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, response *http.Request) {
		fmt.Fprint(writer, "Hi from hi")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, response *http.Request) {
		fmt.Fprint(writer, "Image")
	})

	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, response *http.Request) {
		fmt.Fprint(writer, "Thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
