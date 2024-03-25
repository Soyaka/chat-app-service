package main

import (
	"errors"
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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	connectedUsers map[models.Agent]*websocket.Conn
	mu             sync.Mutex
}

func (s *Server) handleWs(ws *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	//defer ws.Close()

	fmt.Println("new connection from remote addr", ws.RemoteAddr())

	var user models.Agent
	if err := ws.ReadJSON(&user); err != nil {
		fmt.Println("error reading user:", err)
		return
	}

	s.mu.Lock()
	s.connectedUsers[user] = ws
	s.mu.Unlock()

	go s.ReadLoop(ws, user)
}

func (s *Server) ReadLoop(ws *websocket.Conn, user models.Agent) {
	var msg models.Message
	//msg = *models.NewMessage()
	for {
		err := ws.ReadJSON(&msg)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection closed", user)
				break
			}
			fmt.Println("error reading message", err)
			//TODO: here is the error acur  !!!
			panic(err)
		}
		fmt.Println("message:", msg.Content, "from", user)
		msge := models.NewMessage(msg.Content, msg.From, msg.To, "text")
		if err := s.Conversationer(msg.To, msge); err != nil {
			fmt.Println("error sending message to", msg.To, ":", err)
		}
	}
	recover()
}

func (s *Server) Conversationer(to string, msg *models.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	conn, ok := s.getSlectedUser(to)
	if !ok {
		return errors.New("user not found")
	}
	err := conn.WriteJSON(msg)
	if err != nil {
		fmt.Println("error sending message to", to, ":", err)
	}
	return nil
}

func main() {
	var wg sync.WaitGroup
	server := &Server{
		connectedUsers: make(map[models.Agent]*websocket.Conn),
	}
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

func (s *Server) getSlectedUser(to string) (*websocket.Conn, bool) {
	for user, conn := range s.connectedUsers {
		if user.Email == to {
			return conn, true
		}
	}
	return nil, false
}
