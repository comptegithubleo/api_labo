package v1

import (
	"fmt"
	"log"
	"main/utils"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//fetch and receive json from router with all infos. based on user authentication, hide all pending invites except his

	//fetch users
	//fetch pending
	id, err := utils.GetUserId()
	if err != nil {
		log.Println("[X] Error GetUsers: ", err)
		http.Error(w, "Failed to parse your ip to authenticate", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "id is : %d\n", id)
}

func GetPoolMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pool members\n")
}

func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[client] server status\n")
}
