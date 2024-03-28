package handlers

import (
	"encoding/json"
	utils "main/Utils"
	"main/database"
	"main/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Email == "" || user.Password == "" || user.Username == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	result, err := database.InsertUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result.InsertedID.(string)))

}
