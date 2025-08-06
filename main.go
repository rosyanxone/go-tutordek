package main

import (
	"log"
	"net/http"
	"tutordek/handler"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About page"))
	}

	mux.HandleFunc("/", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/about", aboutHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
