package main

import (
	"main/database"
	"main/handlers"
	"main/middleware"
	"main/models"
	"main/websockets"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	server := models.CreateServer()
	client := database.Connect()

	http.HandleFunc("/register", middleware.AuthJWT(func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(w, r, client)
	}))

	http.HandleFunc("/login", middleware.AuthJWT(func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(w, r, client)
	}))

	http.HandleFunc("/refreshToken", middleware.AuthJWT(handlers.RefreshToken))
	http.HandleFunc("/logout", middleware.AuthJWT(handlers.Logout))
	http.HandleFunc("/wsMsg", middleware.AuthJWT(func(w http.ResponseWriter, r *http.Request) {
		websockets.WsMesgHandler(w, r, server, &wg)
	}))
	http.HandleFunc("/wsContact", middleware.AuthJWT(func(w http.ResponseWriter, r *http.Request) {
		websockets.WsContactHandler(w, r, server)
	}))

	http.ListenAndServe(":4444", nil)
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
