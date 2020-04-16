# host-webxr-golang

Hosting WebXR localy using Golang

## Instalation

1) `openssl genrsa -out server.key 2048`
2) `openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`
3) Run! `go run main.go`
