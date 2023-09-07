package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

var UPSTREAM_URL = os.Getenv("UPSTREAM_URL")
var hostname, _ = os.Hostname()

func main() {
	log.Println("Starting proxy...", hostname)
	log.Println("UPSTREAM_URL", UPSTREAM_URL)

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", endpointHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("request received on proxy", r)
	var body io.Reader
	if bodyContent, err := io.ReadAll(r.Body); err == nil && len(bodyContent) > 0 {
		body = bytes.NewBuffer(bodyContent)
	}
	rq, _ := http.NewRequest(http.MethodGet, UPSTREAM_URL, body)
	rs, err := http.DefaultClient.Do(rq)
	log.Println(rs, err)
	if rs != nil {
		w.WriteHeader(rs.StatusCode)
		cnt, _ := io.ReadAll(rs.Body)
		for s := range rs.Header {
			w.Header().Add(s, rs.Header.Get(s))
		}
		w.Header().Add("X-Proxy-Hostname", hostname)
		w.Write(cnt)
	}
}
