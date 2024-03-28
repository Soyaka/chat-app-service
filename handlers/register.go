package handlers

import (
	utils "main/Utils"
	"main/database"
	"main/models"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Email = r.FormValue("email")
	user.Username = r.FormValue("username")
	user.Password, err = utils.HashPassword(r.FormValue("password"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.CreatedAt = time.Now()

	if user.Email == "" || user.Username == "" || user.Password == "" {
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
