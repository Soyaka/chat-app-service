package handlers

import (
	"encoding/json"
	utils "main/Utils"
	"main/database"
	"main/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(w http.ResponseWriter, r *http.Request, client *mongo.Client ) {
	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Email == "" || user.Password == "" {
		http.Error(w, "empty fields", http.StatusBadRequest)
		return
	}

	IntendedUser, err := database.GetUser(bson.M{"email": user.Email}, client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !utils.CheckPasswordHash(user.Password, IntendedUser.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := utils.UserClaims{
		ID: IntendedUser.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	tokenString, err := utils.GenerateJWToken(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	

	w.Write([]byte(tokenString))
}
