package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Use the mux.Handle() function to register the file server as the handler for
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// create router
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// create server that works on 4000
	log.Println("Starting server on 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
