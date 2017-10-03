package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type playroom struct {
	ID             string
	tables         []Table
	players        []Player
	connectionPool []connectionData
}

type connectionData struct {
	ID     string
	name   string
	wsconn *websocket.Conn
}

var pr = playroom{
	ID:             "xxxxxxx",
	tables:         []Table{},
	connectionPool: []connectionData{},
}

func addPlayerToPlayroom(conn *websocket.Conn, msg []byte) string {

	var playerMessage NewPlayerMessage
	err := json.Unmarshal([]byte(msg), &playerMessage)

	if err != nil {
		fmt.Println("Error unmarshalling to NewPlayerMessage", err)
		return "failure"
	}

	var player Player
	player = playerMessage.Payload

	fmt.Printf("Player %s is connected\n", player.Name)

	pr.connectionPool = append(pr.connectionPool, connectionData{"xxxxxxx", player.Name, conn})

	return "success"
}

func addTableToPlayroom(conn *websocket.Conn, msg []byte) string {

	var tableMessage NewTableMessage
	err := json.Unmarshal([]byte(msg), &tableMessage)

	if err != nil {
		fmt.Println("Error unmarshalling to NewTableMessage", err)
		return "failure"
	}

	var table Table
	table = tableMessage.Payload

	pr.addTable(table)

	return "success"
}

func (pr *playroom) addTable(t Table) {
	(*pr).tables = append((*pr).tables, t)

	broadcastmessage := OutgoingMessage{
		Messagetype: "tables",
		Payload:     pr.tables,
	}

	marshalledMessage, err := json.Marshal(broadcastmessage)

	if err != nil {
		fmt.Println("Error marshalling to table struct", err)
	}

	fmt.Printf("\nmarshalledMessage : %s", marshalledMessage)

	fmt.Println("connection pool size", len(pr.connectionPool))

	for _, cd := range pr.connectionPool {
		err := cd.wsconn.WriteMessage(websocket.TextMessage, marshalledMessage)

		if err != nil {
			fmt.Println("Error broadcasting new table to pool", err)
		}
	}
}
