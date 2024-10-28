package main

import (
	"log"
	"net/http"

	"github.com/ldcmleo/blog-api/handlers"
)

func main() {
	mux := http.NewServeMux()

	th := handlers.TestHandler("Hola mundo")
	mux.Handle("/test", th)

	log.Print("Listening... ")
	http.ListenAndServe(":3000", mux)
}
