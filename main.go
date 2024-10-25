package main

import (
	"log"
	"net/http"
)

func testHandler(myVar string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(myVar))
	}
}

func main() {
	mux := http.NewServeMux()

	th := testHandler("hola mundo")
	mux.Handle("/test", th)

	log.Print("Listening... ")
	http.ListenAndServe(":3000", mux)
}
