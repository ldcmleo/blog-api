package main

import (
	"log"
	"net/http"

	"github.com/ldcmleo/blog-api/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/posts", handlers.GetBlogPosts)
	mux.HandleFunc("POST /api/posts", handlers.CreateBlogPost)
	mux.HandleFunc("GET /api/posts/{id}", handlers.GetBlogPost)
	mux.HandleFunc("PUT /api/posts/{id}", handlers.UpdateBlogPost)
	mux.HandleFunc("DELETE /api/posts/{id}", handlers.DeleteBlogPost)

	log.Print("Listening... ")
	http.ListenAndServe(":3000", mux)
}
