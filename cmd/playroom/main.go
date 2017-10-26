package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jar0m1r/playroom"
)

//initialize playroom
var mainroom = playroom.Playroom{}

func main() {

	mainroom = playroom.NewPlayroom()
	go mainroom.BroadcastTableList()

	//initialize mux and listenAndServe

	r := mux.NewRouter()

	//Handle RESTful API calls to User
	u := r.PathPrefix("/api/users").Subrouter()
	u.HandleFunc("/", UsersHandler)
	u.HandleFunc("/{key}/", UserHandler)

	//Handle RESTful API calls to Table
	t := r.PathPrefix("/api/tables").Subrouter()
	t.HandleFunc("/", PostTableHandler).Methods("POST")
	t.HandleFunc("/{key}", GetTableHandler).Methods("GET")
	t.HandleFunc("/{key}/join", joinTableHandler).Methods("POST")

	//Handle WebSocket connections
	w := r.PathPrefix("/ws").Subrouter()
	w.HandleFunc("/", playroomWSServer)
	w.HandleFunc("/{key}", tableWSServer)

	//Handle static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	//ListenAndServe
	fmt.Println("Listening for http on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
