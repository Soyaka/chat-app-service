package websockets

import (
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
	fmt.Println("New connection from remote addr", ws.RemoteAddr())
	var user *models.Agent
	if err := ws.ReadJSON(&user); err != nil {
		fmt.Println("Error reading user:", err)
		return
	}
	fmt.Println("User:", user.Email, user.Username)
	Server.AddUser(user, ws)
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

func ReadLoop(ws *websocket.Conn, Server *models.Server, user *models.Agent) {
	defer Server.RemoveUser(user)
	var msg models.Message
	for {
		if err := ws.ReadJSON(&msg); err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed for user:", user.Email)
			} else {
				fmt.Println("Error reading message:", err)
			}
			break
		}

		msge := models.NewMessage(msg.Content, msg.From, msg.To, "text")
		if err := Conversationer(Server, msg.From, msg.To, msge); err != nil {
			fmt.Println("Error sending message to", msg.To, ":", err)
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

func Conversationer(Server *models.Server, from, to string, msg *models.Message) error {
	Server.Mu.Lock()
	defer Server.Mu.Unlock()
	if err := GetSelectedUsers(Server, msg); err != nil {
		return err
	}
	return nil
}

func GetSelectedUsers(Server *models.Server, msg *models.Message) error {
	for user, conn := range Server.ConnectedUsers {
		if user.Email == msg.To || user.Email == msg.From {
			err := WriteMessage(conn, msg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func WriteMessage(conn *websocket.Conn, msg *models.Message) error {
	err := conn.WriteJSON(msg)
	if err != nil {
		return err
	}
	return nil
}

/* #####WRITE UNIT TESTS####*/
