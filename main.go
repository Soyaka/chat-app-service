package main

import (
	"main/database"
	"main/handlers"
	"main/middleware"
	"main/models"
	"main/websockets"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/x/mongo/driver/auth"
)

func main() {
	var wg sync.WaitGroup
	server := models.CreateServer()
	client := database.Connect()
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	router := r.PathPrefix("/api").Subrouter()
	router.Use(mux.CORSMethodMiddleware(r))

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(w, r, client)
	}).Methods("POST")
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(w, r, client)
	}).Methods("POST")
	router.HandleFunc("/refreshToken", middleware.AuthJWT(handlers.RefreshToken)).Methods("POST")
	router.HandleFunc("/logout", middleware.AuthJWT(handlers.Logout)).Methods("POST")
	router.HandleFunc("/wsmsg", middleware.AuthJWT(func(w http.ResponseWriter, r *http.Request) {
		websockets.WsMesgHandler(w, r, server, &wg)
	})).Methods("GET", "OPTIONS")
	router.HandleFunc("/wscontact", middleware.AuthJWT(func(w http.ResponseWriter, r *http.Request) {
		websockets.WsContactHandler(w, r, server)
	})).Methods("GET", "OPTIONS")

	handler := c.Handler(r)
	http.ListenAndServe(":4444", handler)
	wg.Wait()
}

//FIXME: Add the error handling
//FIXME: handle db client distribution between handlers

/* TODO:
#Fix the TODOS
#Add Group Chat
#Add Authentication
#Add Encryption
#Add database integration
#Add redis
#Add server handling in case a lot of users and in case of server crash or restart : add replicas
*/
