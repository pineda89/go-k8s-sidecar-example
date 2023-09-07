package main

import (
	"log"
	"net/http"
	"os"
)

var SERVICE_PORT = getenvOrDefault("SERVICE_PORT", "3000")
var hostname, _ = os.Hostname()

func getenvOrDefault(s string, s2 string) string {
	if v := os.Getenv(s); len(v) > 0 {
		return v
	}
	return s2
}

func main() {
	log.Println("Starting service...")

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", endpointHandler)
	log.Fatalln(http.ListenAndServe(":"+SERVICE_PORT, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("request received", r)
	w.Header().Add("X-Service-Hostname", hostname)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world from the service"))
}
