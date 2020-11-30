package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Punishment struct {
	UUID        string `json:"UUID"`
	Username    string `json:"username"`
	ExpiryDate  string `json:"expiryDate"`
	Offense     string `json:"offense"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

var Punishments []Punishment
var myRouter = mux.NewRouter().StrictSlash(true)

func returnAllPunishments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPunishments")
	v := r.URL.Query()
	pageSize := v.Get("pageSize")
	page := v.Get("page")
	fmt.Println(pageSize)
	fmt.Println(page)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Punishments)
}

func returnSingularPunisment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["UUID"]

	for _, punishment := range Punishments {
		if punishment.UUID == key {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(punishment)
		}
	}
}

func handleRequests() {
	myRouter.HandleFunc("/allPunishments", returnAllPunishments)
	myRouter.HandleFunc("/punishment/{uuid}", returnSingularPunisment)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Punishments = []Punishment{
		Punishment{UUID: "257a1a94-7dc1-4657-92e3-22b2f1ea7573", Username: "ocelotcr", ExpiryDate: "10000000000", Offense: "Hacking", Description: "I saw the boi hacking", Type: "PERMA"},
	}
	handleRequests()
}
