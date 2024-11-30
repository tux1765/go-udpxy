package main

import (
	"github.com/tux1765/go-udpxy/udpxy"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/stream", udpxy.StreamFromUdp)
	log.Println("Serving MPEG-TS UDP stream on :8080/stream")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
