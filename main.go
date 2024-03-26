package main

import (
	"main/models"
	"main/websockets"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	server := models.CreateServer()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websockets.WsHandler(w, r, server, &wg)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
			
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		
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