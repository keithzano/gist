package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This is the home page</h1>")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	fmt.Println("The server is starting on port :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}

}
