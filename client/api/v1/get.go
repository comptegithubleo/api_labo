package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/utils"
	"net/http"
)

type PendingInvites struct {
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
	response, err := http.Get("http://localhost:3032/v1/invites") //hardcoded for now
	if err != nil {
		http.Error(w, "Failed to retrieve invites", http.StatusInternalServerError)
		log.Println("[X] Error GetInvites http.Get: ", err)
		return
	}
	defer response.Body.Close()

	invitesResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var invites []PendingInvites
	json.Unmarshal(invitesResponse, &invites)

	//from his ip, we get his id, hide all pending invites that's not his
	user_id, err := utils.GetUserId()
	if err != nil {
		log.Println("[X] Error GetInvites GetUserId: ", err)
		http.Error(w, "Failed to parse your ip to authenticate", http.StatusBadRequest)
		return
	}
	var filteredInvites []PendingInvites
	for i := 0; i < len(invites); i++ {
		if invites[i].From == user_id || invites[i].To == user_id {
			filteredInvites = append(filteredInvites, invites[i])
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(filteredInvites); err != nil {
		http.Error(w, "Failed to retrieve invites", http.StatusInternalServerError)
		log.Println("[X] Error GetInvites json.Encode: ", err)
		return
	}
}

// Deprecated because available in GetUsers data ?

/* func GetPoolMembers(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetUserId()
	if err != nil {
		log.Println("[X] Error GetPoolMembers GetUserId: ", err)
		http.Error(w, "Failed to parse your ip to authenticate", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "id is : %d\n", id)
	fmt.Fprintf(w, "pool members\n")
}
 */
func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[client] server status\n")
}
