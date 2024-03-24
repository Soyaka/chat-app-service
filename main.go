package main

import (
	"fmt"
	"io"
	"net/http"
	"sockets/models"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type User *models.Agent
type Server struct {
	connectedUsers map[User]bool
}

func NewServer() *Server {
	return &Server{
		connectedUsers: make(map[User]bool),
	}
}

func (S *Server) handleWs(ws *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("new connection from remote addr", ws.RemoteAddr())
	S.ReadLoop(ws)
	user := models.NewUser(ws)
	S.connectedUsers[user] = true

}

func (s *Server) ReadLoop(ws *websocket.Conn) {
	var msg models.Message
	for {
		err := ws.ReadJSON(&msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error", err)
			continue
		}
		fmt.Println("message:", msg.Content, "from", ws.RemoteAddr())
		msge := models.NewMessage("hello eve", msg.From, msg.To, "text")
		ws.WriteJSON(msge)
	}
}

func Conversationer(users []User, to User, msg *models.Message) {
	for _, user := range users {
		if user == to {
			err := user.ChatConn.WriteJSON(msg)
			if err != nil {
				fmt.Println("err sending msg ", err)
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	server := NewServer()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		wg.Add(1)
		go server.handleWs(conn, &wg)
	})
	http.ListenAndServe(":4444", nil)
	wg.Wait()
}
