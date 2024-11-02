package main

import (
	"log"
	"net/http"

	"github.com/ldcmleo/blog-api/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", handlers.GetBlogPosts())

	log.Print("Listening... ")
	http.ListenAndServe(":3000", mux)
}
