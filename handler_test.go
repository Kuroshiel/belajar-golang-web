package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// Handler Golang Web

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "Localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// ServeMux Golang Web

func TestServeMux(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "image")
	})

	mux.HandleFunc("/images/thumbsnail/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Thumbsnail")
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

// Request Web Golang

func TestRequesr(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "Localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
