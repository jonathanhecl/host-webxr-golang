package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Jonathan G. Hecl || www.jonathanhecl.com || 2020

// 1) openssl genrsa -out server.key 2048
// 2) openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
// 3) Run! go run main.go

func main() {
	http.HandleFunc("/", fileServer)
	fmt.Println("Hosting WebXR on https://localhost:443")
	if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil); err != nil {
		panic(err)
	}
}

func fileServer(w http.ResponseWriter, r *http.Request) {
	ruri := r.RequestURI
	if strings.HasSuffix(ruri, ".mem") || strings.HasSuffix(ruri, ".data") {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.FileServer(http.Dir("./webxr")).ServeHTTP(w, r)
}
