package playroom

import "github.com/gorilla/websocket"

//User is a playroom type User
type User struct {
	ID          string          `json:"id,omitempty"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	Credit      float32         `json:"credit,omitempty"`
	Wsconn      *websocket.Conn `json:"-"`
	Wstableconn *websocket.Conn `json:"-"`
}
