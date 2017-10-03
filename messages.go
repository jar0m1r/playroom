package main

/* INCOMING MESSAGES */

//IncomingMessage is a flexible struct for receiving incomming message
type IncomingMessage struct {
	Messagetype string
	Payload     interface{}
}

//NewPlayerMessage is a concrete struct for a New Player Message
type NewPlayerMessage struct {
	Messagetype string
	Payload     Player
}

//NewTableMessage is a concrete struct for a New Table Message
type NewTableMessage struct {
	Messagetype string
	Payload     Table
}

//Player is a concrete struct for Player
type Player struct {
	ID   string
	Name string
}

//Table is a concrete struct for table
type Table struct {
	ID         string
	Name       string
	Maxplayers int
	Players    []Player
}

/* OUTGOING MESSAGES */

//OutgoingMessage is used to broadcast messages to the client
type OutgoingMessage struct {
	Messagetype string
	Payload     interface{}
}

//ResultMessage is used to send a success/failure to the client
type ResultMessage struct {
	Result      string //success or failure
	Description string
}
