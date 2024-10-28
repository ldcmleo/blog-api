package handlers

import (
	"net/http"
)

func TestHandler(myVar string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(myVar))
	}
}
