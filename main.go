package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "<h1>This is the home page</h1>")
}

func gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific gist with ID : %d", id)

}

func gistCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("This is The create page"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/view", gistView)
	mux.HandleFunc("/create", gistCreate)
	fmt.Println("The server is starting on port :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}

}
