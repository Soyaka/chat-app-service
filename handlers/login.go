package handlers

import (
	utils "main/Utils"
	"main/database"
	"main/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	if user.Email == "" || user.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	IntendedUser, err := database.GetUser(bson.M{"email": user.Email})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !utils.CheckPasswordHash(user.Password, IntendedUser.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	//TODO:Generate JWT Token

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}