package websockets

import (
	"fmt"
	"main/models"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{

	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsMesgHandler(w http.ResponseWriter, r *http.Request, S *models.Server, wg *sync.WaitGroup) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	wg.Add(1)
	go HandleConnections(conn, S, wg)

}

func WsContactHandler(w http.ResponseWriter, r *http.Request, S *models.Server) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	contacts := S.GetConnectedUsers()
	nbRows := len(contacts)
	contact := &ConectedUsers{
		data:   contacts,
		nbRows: nbRows,
	}
	if err = conn.WriteJSON(contact); err != nil {
		fmt.Println("error in #wsHandlers:f:WsContactHandler")
		panic(err)
	}

}

type ConectedUsers struct {
	data   []*models.Agent
	nbRows int
}
