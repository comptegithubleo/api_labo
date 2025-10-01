package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/utils"
	"net/http"
)

type PendingConnection struct {
	From int `json:"from"`
	To   int `json:"to"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:3032/v1/users") //hardcoded for now
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		log.Println("[X] Error GetUsers http.Get: ", err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(w, response.Body)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		log.Println("[X] Error GetUsers io.Copy: ", err)
		return
	}
}

func GetInvites(w http.ResponseWriter, r *http.Request) {
	//based on user authentication (ip), hide all pending invites except his
	response, err := http.Get("http://localhost:3032/v1/invites") //hardcoded for now
	if err != nil {
		http.Error(w, "Failed to retrieve invites", http.StatusInternalServerError)
		log.Println("[X] Error GetUsers http.Get: ", err)
		return
	}
	defer response.Body.Close()

	invitesResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var invites []PendingConnection
	json.Unmarshal(invitesResponse, &invites)
	
	log.Println(invitesResponse)
	for i := 0; i < len(invites); i++ {
		log.Println(invites[i].From)
	}
}

func GetPoolMembers(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetUserId()
	if err != nil {
		log.Println("[X] Error GetUsers: ", err)
		http.Error(w, "Failed to parse your ip to authenticate", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "id is : %d\n", id)
	fmt.Fprintf(w, "pool members\n")
}

func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[client] server status\n")
}
