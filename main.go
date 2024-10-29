package main

import (
	"context"
	"log"

	"github.com/ldcmleo/blog-api/database"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	db, err := database.GetDatabaseClient()
	if err != nil {
		panic("error getting database" + err.Error())
	}

	dbNames, nmErr := db.ListDatabaseNames(context.TODO(), bson.D{})
	if nmErr != nil {
		panic("error nmerr: " + nmErr.Error())
	}

	for _, names := range dbNames {
		log.Println(names)
	}

	// mux := http.NewServeMux()

	// th := handlers.TestHandler("Hola mundo")
	// mux.Handle("/test", th)

	// log.Print("Listening... ")
	// http.ListenAndServe(":3000", mux)
}
