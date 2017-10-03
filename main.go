package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.HandleFunc("/api/login", handleLogin)

	http.HandleFunc("/websocket", socketServer)

	fmt.Println("Listening for http on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var player Player
	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		err = json.Unmarshal([]byte(body), &player)

		if err != nil {
			fmt.Println("Error unmarshalling login json", err)
		}

		player.ID = uuid.New().String()

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(player)
}

func socketServer(w http.ResponseWriter, r *http.Request) {

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

			addPlayerToPlayroom(conn, msg)

			/* 			err = conn.WriteMessage(msgType, []byte(result))

			   			if err != nil {
			   				fmt.Println(err)
			   				return
			   			} */

		} else if incomingMessage.Messagetype == "table" {

			addTableToPlayroom(conn, msg)

			/* 			err = conn.WriteMessage(msgType, []byte(result))

			   			if err != nil {
			   				fmt.Println(err)
			   				return
			   			} */

		} else {
			conn.Close()
			fmt.Println("Connection closed by server on last message:", msg)
			return
		}
	}

}
