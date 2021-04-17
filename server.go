package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "%s	%s	%s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Header:\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "\t%s:\n\t\t%s\n", k, v)
	}
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "Remote: %s\n", r.RemoteAddr)

	fmt.Fprintf(w, "URL.Path: %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}

