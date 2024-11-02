package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ldcmleo/blog-api/db"
	"github.com/ldcmleo/blog-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Disconnect(ctx)

	collection := db.Database("myblog").Collection("blogposts")

	filter := bson.D{}

	result, resErr := collection.Find(ctx, filter)
	if resErr != nil {
		http.Error(w, "Error al procesar los datos de la coleccion de blogposts", http.StatusInternalServerError)
		return
	}

	defer result.Close(ctx)

	var blogPosts []models.BlogPost
	if err := result.All(ctx, &blogPosts); err != nil {
		http.Error(w, "Error al decodificar los datos", http.StatusInternalServerError)
		return
	}

	if len(blogPosts) == 0 {
		w.Write([]byte("[]"))
		return
	}

	if err := json.NewEncoder(w).Encode(blogPosts); err != nil {
		http.Error(w, "Error al enviar los datos", http.StatusInternalServerError)
	}
}

func CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.BlogPost

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Error, datos no validos", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Error accediendo a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Disconnect(ctx)

	collection := db.Database("myblog").Collection("blogposts")

	// modificaciones al post
	post.CreatedAt = time.Now()

	result, err := collection.InsertOne(ctx, post)
	if err != nil {
		http.Error(w, "Error al guardar los datos", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error al retornar los datos", http.StatusInternalServerError)
		return
	}
}

func GetBlogPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	postID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "error al procesar este ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Disconnect(ctx)

	collection := db.Database("myblog").Collection("blogposts")
	filter := bson.M{"_id": postID}

	var post models.BlogPost

	if err := collection.FindOne(ctx, filter).Decode(&post); err != nil {
		http.Error(w, "Error al decodificar la informacion", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Error al mostrar los datos", http.StatusInternalServerError)
		return
	}
}

func UpdateBlogPost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	postID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Error al procesar ID", http.StatusBadRequest)
		return
	}

	var post models.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Error al procesar la data", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Disconnect(ctx)

	collection := db.Database("myblog").Collection("blogposts")
	filter := bson.M{"_id": postID}

	update := bson.M{
		"$set": bson.M{
			"title":   post.Title,
			"content": post.Content,
			"tags":    post.Tags,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	if err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&post); err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Error documento no encontrado", http.StatusNotFound)
			return
		}

		http.Error(w, "Error al actualizar los datos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Error al mostrar los datos", http.StatusInternalServerError)
		return
	}
}

func DeleteBlogPost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	postID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Error al procesar ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := db.Connect()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}
	defer db.Disconnect(ctx)

	collection := db.Database("myblog").Collection("blogposts")
	filter := bson.M{"_id": postID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		http.Error(w, "Error al eliminar los datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result.DeletedCount); err != nil {
		http.Error(w, "Error al mostrar los datos", http.StatusInternalServerError)
		return
	}
}
