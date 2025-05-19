package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/view", view)
	mux.HandleFunc("/create", create)

	fmt.Printf("Staring Server on port: %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
