package main

import (
	"main/handlers"
	"main/models"
	"main/websockets"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	server := models.CreateServer()
	http.HandleFunc("/wsMsg", func(w http.ResponseWriter, r *http.Request) {
		websockets.WsMesgHandler(w, r, server, &wg)
	})
	http.HandleFunc("/wsContact", func(w http.ResponseWriter, r *http.Request) {
		websockets.WsContactHandler(w, r, server)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(w, r)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(w, r)
	})

	http.HandleFunc("/refreshToken", func(w http.ResponseWriter, r *http.Request) {
		handlers.Refresh(w, r)

	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		handlers.Logout(w, r)
	})

	http.ListenAndServe(":4444", nil)
	wg.Wait()
}



/* TODO:
#Fix the TODOS
#Add Group Chat
#Add Authentication
#Add Encryption
#Add database integration
#Add redis
#Add server handling in case a lot of users and in case of server crash or restart : add replicas
*/
