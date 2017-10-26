package main

/* INCOMING MESSAGES */

//IncomingMessage is a flexible struct for receiving incomming message
type IncomingMessage struct {
	Messagetype string      `json:"messagetype"`
	Payload     interface{} `json:"payload"`
}

//NewUserMessage struct for a New User Messages
type NewUserMessage struct {
	Messagetype string      `json:"messagetype"`
	Payload     UserMessage `json:"payload"`
}

//NewTableMessage struct for a New Table Messages
type NewTableMessage struct {
	Messagetype string       `json:"messagetype"`
	Payload     TableMessage `json:"payload"`
}

//UserMessage struct for User Messages
type UserMessage struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name"`
	Credit float32 `json:"credit"`
}

//TableMessage struct for Table Messages
type TableMessage struct {
	ID       string        `json:"id,omitempty"`
	Name     string        `json:"name"`
	Gametype string        `json:"gametype"`
	Numseats int           `json:"numseats"`
	Users    []UserMessage `json:"users,omitempty"`
}

/* OUTGOING MESSAGES */

//OutgoingMessage is used to broadcast messages to the client
type OutgoingMessage struct {
	Messagetype string      `json:"messagetype"`
	Payload     interface{} `json:"payload"`
}

//ResultMessage is used to send a success/failure and a description to the client
type ResultMessage struct {
	Result      string `json:"result"` //success or failure
	Description string `json:"description"`
}
