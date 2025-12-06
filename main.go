package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port string

func init() {
	flag.StringVar(&port, "port", "4000", "Port to listen on")
}

func main() {
	flag.Parse()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	fmt.Println("Listening on port", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
