package handlers

import (
	"encoding/json"
	"log"
	utils "main/Utils"
	"main/database"
	"main/models"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func Register(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println(" error  decoding user", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Email == "" || user.Password == "" || user.Username == "" {

		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Println(" error hashing pwd:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := models.NewUser(user.Username, user.Email, user.Password)

	result, err := database.InsertUser(newUser, client)
	if err != nil {
		log.Println(" error  inserting in database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"insertedID": result.InsertedID})

}


//Tested successfully