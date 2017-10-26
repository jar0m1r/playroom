package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/websocket"
	"github.com/jar0m1r/playroom"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func playroomWSServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receiving sebsocket message..")
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()

		fmt.Printf("Received a message %s", msg)

		if err != nil {
			fmt.Println("ReadMessage error:", err)
			return
		}

		var incomingMessage IncomingMessage

		err = json.Unmarshal([]byte(msg), &incomingMessage)

		if err != nil {
			fmt.Println("Error unmarshalling to incommingMessage", err)
			return
		}

		if incomingMessage.Messagetype == "person" {

			var nu NewUserMessage
			err := json.Unmarshal([]byte(msg), &nu)

			if err != nil {
				fmt.Printf("Error unmarshalling userMessage %s", err)
			}

			var u playroom.User
			u.ID = nu.Payload.ID
			u.Name = nu.Payload.Name

			err = mainroom.AddUserConnection(conn, u)

		} else {
			conn.Close()
			fmt.Println("Connection closed by server on last message:", msg)
			return
		}
	}

}

func tableWSServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tableid := vars["key"]
	fmt.Printf("Receiving websocket message for table %s", tableid)
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()

		fmt.Printf("Received a message %s", msg)

		if err != nil {
			fmt.Println("ReadMessage error:", err)
			return
		}

		var incomingMessage IncomingMessage

		err = json.Unmarshal([]byte(msg), &incomingMessage)

		if err != nil {
			fmt.Println("Error unmarshalling to incommingMessage", err)
			return
		}

		if incomingMessage.Messagetype == "person" {

			var nu NewUserMessage
			err := json.Unmarshal([]byte(msg), &nu)

			if err != nil {
				fmt.Printf("Error unmarshalling userMessage %s", err)
			}

			//get user from table userMap, add user connection to tableBroadcaster
			if table, ok := mainroom.TableMap[tableid]; ok {
				if user, ok := table.UserMap[nu.Payload.ID]; ok {
					(*user).Wstableconn = conn
				}
			}

		} else {
			conn.Close()
			fmt.Println("Connection closed by server on last message:", msg)
			return
		}
	}
}
