package controllers

import (
	"blog-api/config"
	"blog-api/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	config.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	config.DB.Create(&post)
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		http.NotFound(w, r)
		return
	}
	json.NewDecoder(r.Body).Decode(&post)
	config.DB.Save(&post)
	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	config.DB.Delete(&models.Post{}, id)
	w.WriteHeader(http.StatusNoContent)
}
