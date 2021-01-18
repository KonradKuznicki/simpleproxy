package main

import (
	"fmt"
	"github.com/armon/go-socks5"
	"net/http"
)

func main() {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, "{\"status\": \"ok\"}\n")
		})

		http.ListenAndServe("0.0.0.0:8001", nil)
	}()

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "0.0.0.0:8000"); err != nil {
		panic(err)
	}
}
