package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "http.MethodGet")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello There!")
}

func form(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post Request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s", name)
	fmt.Fprintf(w, "Address = %s", address)
}

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/", fileServer)

	mux.HandleFunc("/hello", hello)

	mux.HandleFunc("/form", form)

	log.Print("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}

}
