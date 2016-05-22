package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "7777", "Listen port.")
	path = flag.String("path", ".", "Path dir.")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.Printf("Listen [%v]", *port)
	log.Printf("Path [%v]", *path)
	http.Handle("/", http.FileServer(http.Dir(*path)))
	http.ListenAndServe(":"+*port, nil)
	log.Printf("Server exit")
}
