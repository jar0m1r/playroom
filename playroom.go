package playroom

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jar0m1r/blackjack"
)

//Playroom is the core Playroom Struct for use in servers to hold Users, Tables and Connections
type Playroom struct {
	ID       string
	Tables   []Table
	TableMap map[string]*Table
	UserMap  map[string]*User
}

//NewPlayroom initializes and returns a new Playroom struct
func NewPlayroom() Playroom {
	p := Playroom{
		ID:       "xxxxxxx",
		Tables:   []Table{},
		TableMap: map[string]*Table{},
		UserMap:  map[string]*User{},
	}

	return p
}

//AddUser creates a (Playroom) User with uuid and credit and adds to the Playroom Users slice
func (pr *Playroom) AddUser(name string) (User, error) {

	user := User{
		ID:     uuid.New().String(),
		Name:   name,
		Credit: 10.0,
		Wsconn: nil,
	}

	(*pr).UserMap[user.ID] = &user

	fmt.Println("User created with id", user.ID)

	return user, nil //Todo Error handling
}

//AddUserConnection reads the user and adds the user and connection to the ConnectionsPool
func (pr *Playroom) AddUserConnection(conn *websocket.Conn, u User) error {

	user := (*pr).UserMap[u.ID]
	(*user).Wsconn = conn
	fmt.Println("connection added to ", *user)
	return nil
}

//AddTable creates a table of gametype and returns the table ID
func (pr *Playroom) AddTable(name, gametype string, numseats int) (string, error) {

	t, err := NewTable(name, gametype, numseats)

	if err != nil {
		return "", fmt.Errorf("Table could not be created %s", err)
	}

	(*pr).Tables = append((*pr).Tables, t)
	(*pr).TableMap[t.ID] = &t

	//start broadcasting table states from the moment it is created
	t.BroadcastTable()

	return t.ID, nil
}

//BroadcastTableList is a temporary solution for broadcasting a table list via websockets when changed
func (pr *Playroom) BroadcastTableList() {
	var prevlen int

	for {
		time.Sleep(time.Second * 10)
		if len(pr.Tables) != prevlen {
			prevlen = len(pr.Tables)
			tables, err := json.Marshal(pr.Tables)
			fmt.Printf("tables after marshall %s\n", tables)
			if err != nil {
				fmt.Println("Error marshalling tables", err)
			}

			for _, v := range pr.UserMap {
				err := v.Wsconn.WriteMessage(websocket.TextMessage, tables)

				if err != nil {
					fmt.Println(err)
					return
				}

			}
		}
	}
}

//GetTable returns table struct value
func (pr *Playroom) GetTable(key string) (Table, error) {

	if t, ok := (*pr).TableMap[key]; ok {
		return *t, nil
	}
	return Table{}, fmt.Errorf("the key did not correspond to a known table")
}

/* func (pr *Playroom) addTable(t Table) {
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
} */

func startGame(t *blackjack.Table) {
	/* 	(*t).NewGame()

	   	(*t).Deal()

	   	(*t).PrintHands() //one dealer hand blind (later)

	   	for playerindex := range (*t).Players {
	   		(*t).Play(playerindex)
	   	}

	   	(*t).PrintHands() //one dealer hand blind (later)

	   	//check if deck actually schrunk

	   	(*t).FinishGame() */

}
