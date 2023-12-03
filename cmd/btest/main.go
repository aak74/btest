package main

import (
	"net/http"
)

var (
	isStopping = false
)

/*
	func ping(w http.ResponseWriter, req *http.Request) {
		if isStopping {
			w.Header().Set("Rs-Weight", "0")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Stopping\n")
			return
		}
		fmt.Fprintf(w, "OK\n")
	}

	func hello(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n")
	}

	func restart(w http.ResponseWriter, req *http.Request) {
		isStopping = false
		fmt.Fprintf(w, "Starting\n")
	}

	func shutdown(w http.ResponseWriter, req *http.Request) {
		isStopping = true
		fmt.Fprintf(w, "Stopping\n")
	}
*/
func main() {
	// http.HandleFunc("/ping", ping)
	// http.HandleFunc("/hello", hello)
	// http.HandleFunc("/restart", restart)
	// http.HandleFunc("/shutdown", shutdown)
	// http.ListenAndServe(":8090", nil)
	srv := NewServer()

	http.ListenAndServe(":8090", srv)
}
