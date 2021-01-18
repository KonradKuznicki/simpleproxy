package main

import (
	"crypto/tls"
	"fmt"
	"github.com/armon/go-socks5"
	"net/http"
	"time"
)

func checkServer(url string) error {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	_, err := client.Get(url)

	return err
}

func main() {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if err := checkServer("https://10.8.5.37:8443/web/guest/home"); err != nil {
				fmt.Fprint(w, "{\"status\": \"error\", \"error\": \"tunnel down or host not responding, check logs\"}\n")
				fmt.Println(err)
			} else {
				fmt.Fprint(w, "{\"status\": \"ok\"}\n")
			}
		})

		http.ListenAndServe("0.0.0.0:8001", nil)
	}()

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "0.0.0.0:8000"); err != nil {
		panic(err)
	}
}
