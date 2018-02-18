// Package main contains a static server for situations where
// you have no time to waste writing a server by yourself.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.Int("port", 7777, "Listen port.")
	path = flag.String("path", ".", "Path dir.")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.Printf("Listen [%d]", *port)
	log.Printf("Path [%v]", *path)
	http.Handle("/", http.FileServer(http.Dir(*path)))
	addrs := fmt.Sprintf(":%d", *port)
	if err := http.ListenAndServe(addrs, nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("Server exit")
}
