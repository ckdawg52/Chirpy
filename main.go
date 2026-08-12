package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	fmt.Println("Starting Server . . .")

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	s := &http.Server{
		Addr: ":8080",
		Handler: mux, 
	}
	log.Fatal(s.ListenAndServe())
}