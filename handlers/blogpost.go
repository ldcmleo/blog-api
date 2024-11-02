package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ldcmleo/blog-api/db"
	"github.com/ldcmleo/blog-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBlogPosts() http.HandlerFunc {
	db, err := db.Connect()
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Database("myblog").Collection("blogposts")

	filter := bson.D{}

	result, resErr := collection.Find(ctx, filter)
	if resErr != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Error al procesar los datos de la coleccion de blogposts", http.StatusInternalServerError)
		}
	}

	defer result.Close(ctx)

	var blogPosts []models.BlogPost
	if err := result.All(ctx, &blogPosts); err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Error al decodificar los datos", http.StatusInternalServerError)
		}
	}

	if len(blogPosts) < 1 {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			response := map[string]interface{}{}
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Error al enviar los datos", http.StatusInternalServerError)
			}
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(blogPosts); err != nil {
			http.Error(w, "Error al enviar los datos", http.StatusInternalServerError)
		}
	}
}
