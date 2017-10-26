package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//UsersHandler (POST only) handles login
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var u UserMessage
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		err = json.Unmarshal([]byte(body), &u)

		if err != nil {
			fmt.Println("Error unmarshalling login json", err)
		}

		roomUser, err := mainroom.AddUser(u.Name)

		if err != nil {
			fmt.Println("Error adding user to playroom", err)
		}

		u.ID = roomUser.ID

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(u)

	} else {

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)

	}
}

//UserHandler (GET, POST, PUT) handles user C, R, U
func UserHandler(w http.ResponseWriter, r *http.Request) {
	return
}

//PostTableHandler POST (PUT?) handles table Create(Update?)
func PostTableHandler(w http.ResponseWriter, r *http.Request) {
	/* if r.Method == "POST" { */
	var table TableMessage
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	err = json.Unmarshal([]byte(body), &table)

	if err != nil {
		fmt.Println("Error unmarshalling login json", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	id, err := mainroom.AddTable(table.Name, table.Gametype, table.Numseats) //hardcoded numseats, fix

	if err != nil {
		http.Error(w, "Error creating table", http.StatusUnprocessableEntity)
	}

	result := TableMessage{
		ID: id,
	}

	json.NewEncoder(w).Encode(result)

	/* 	}else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	} */
}

//GetTableHandler based on {key} returns Table struct
func GetTableHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	table, err := mainroom.GetTable(key)

	if err != nil {
		http.Error(w, "Not able to return table based on key", http.StatusMethodNotAllowed)
	}

	json.NewEncoder(w).Encode(table)
}

//joinTableHandler adds the user to the table (initially as viewer)
func joinTableHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	//change to using the header token to identify user and make it GET
	var u UserMessage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "not able to read body from request", http.StatusBadRequest)
	}
	fmt.Println("Body read from websocket join request", string(body))

	err = json.Unmarshal(body, &u)
	if err != nil {
		fmt.Println("Error marshalling user from req body", err)
		http.Error(w, "not able to marshall user from request body", http.StatusBadRequest)
	}

	if t, ok := mainroom.TableMap[key]; ok {
		if usr, ok := mainroom.UserMap[u.ID]; ok {
			t.AddUser(usr) //change to user reference later
			json.NewEncoder(w).Encode(t)

		} else {
			http.Error(w, "user not known in playroom", http.StatusBadRequest)
		}

	} else {
		http.Error(w, "not able to identify/marshall table from request body", http.StatusBadRequest)
	}

}
