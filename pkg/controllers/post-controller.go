package controllers

import (
	"encoding/json"
	"go-scristi/pkg/models"
	"net/http"
)

var NewPost models.Post

func GetPosts(w http.ResponseWriter, r *http.Request) {
	newPosts := models.ReturnLastPosts(10)
	res, _ := json.Marshal(newPosts)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
