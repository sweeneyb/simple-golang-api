package main

// shamelessly copied from https://gobyexample.com/http-servers

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func version(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "0.0.1\n")
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/version", version)

	http.ListenAndServe(":8090", nil)
}
