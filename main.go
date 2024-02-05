package main

import (
	"fmt"
	"net/http"
	"time"
)

func init() {
	// sets up global context & is executed before main()
	fmt.Println("Setting up Global Context")
}

func main() {
	fmt.Println("Application Initialised " + time.Now().UTC().String())
	startServer()
}

func startServer() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
	})

	// open a TCP port on 8080 and serve a handler to handle requests on incoming connections - start a HTTP 1.1 server.
	http.ListenAndServe(":8080", nil)
}
