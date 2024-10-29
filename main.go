package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ldcmleo/blog-api/db"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client, err := db.Connect()
	if err != nil {
		log.Fatal("error with database: " + err.Error())
	}

	dbNames, qErr := client.ListDatabaseNames(context.TODO(), bson.D{})
	if qErr != nil {
		panic("error getting database names: " + qErr.Error())
	}

	for _, names := range dbNames {
		fmt.Println(names)
	}

	// mux := http.NewServeMux()

	// th := handlers.TestHandler("Hola mundo")
	// mux.Handle("/test", th)

	// log.Print("Listening... ")
	// http.ListenAndServe(":3000", mux)
}
