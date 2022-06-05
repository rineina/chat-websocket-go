package main

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WebsocketHandler)
}

func main(){
	log.Fatal(http.ListenAndServe("192.168.0.101:8000", nil))
}