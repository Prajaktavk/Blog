package main

import (
	"log"
	"net/http"

	"blog-api/config"
	"blog-api/controllers"
	"blog-api/models"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Post{}) // if using a Post model

	router := mux.NewRouter()

	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")

	log.Println("🚀 Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
