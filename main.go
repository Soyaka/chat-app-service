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

	http.ListenAndServe(":4444", nil)
	wg.Wait()
}
