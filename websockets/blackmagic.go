package websockets

import (
	"errors"
	"fmt"
	"io"
	"main/models"
	"sync"

	"github.com/gorilla/websocket"
)

/* Take the connections and handle them concurrently

# recieve the first letter set a new user
# and pass to the next step for recieving messages
TODO: Add a safe way to recieve messages
TODO: Add a channel for recieving errors that my acur

*/

func HandleConnections(ws *websocket.Conn, Server *models.Server, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("new connection from remote addr", ws.RemoteAddr())
	var user models.Agent
	if err := ws.ReadJSON(&user); err != nil {
		fmt.Println("error reading user:", err)
		panic(err)
	}
	Server.Mu.Lock()
	Server.ConnectedUsers[user] = ws
	Server.Mu.Unlock()
	go ReadLoop(ws, Server, user)

}

/*Loop over the connection and read the messages
#recives messages from the connection
#create a new message and recend it to the next step
TODO: Add a safe way to recieve messages
TODO: Add a channel for recieving errors that my acur
TODO: handle reads concurently
TODO:Take the user base on the ID not email

*/

func ReadLoop(ws *websocket.Conn, Server *models.Server, user models.Agent) {
	var msg models.Message
	for {
		err := ws.ReadJSON(&msg)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection closed", user)
				break
			}
			fmt.Println("error reading message", err)
			panic(err)
		}
		fmt.Println("message:", msg.Content, "from", user)
		msge := models.NewMessage(msg.Content, msg.From, msg.To, "text")
		if err := Conversationer(Server, msg.To, msge); err != nil {
			fmt.Println("error sending message to", msg.To, ":", err)
		}
	}
}

/*Find the reciever User and Reseng the messgae to it

TODO: Add a safe way handle users not found
TODO: handle errors if acur in the step of sending message
TODO: Add onather pahse of sending message
TODO: Add checks if the user is online
TODO: Add entry to store messages in database or redis
TODO:Send message to sender and reciever

*/

func Conversationer(Server *models.Server, to string, msg *models.Message) error {
	Server.Mu.Lock()
	defer Server.Mu.Unlock()
	conn, ok := GetSlectedUser(Server.ConnectedUsers, to)
	if !ok {
		return errors.New("user not found")
	}
	err := conn.WriteJSON(msg)
	if err != nil {
		fmt.Println("error sending message to", to, ":", err)
	}
	return nil
}

func GetSlectedUser(ConnectedUsers map[models.Agent]*websocket.Conn, to string) (*websocket.Conn, bool) {
	for user, conn := range ConnectedUsers {
		if user.Email == to {
			return conn, true
		}
	}
	return nil, false
}
