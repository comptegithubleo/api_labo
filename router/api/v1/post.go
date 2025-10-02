package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"main/utils"
	"net/http"
	"net/url"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create user\n")
}

// if 12 adds 13 :
//
//	if invites contains "from 13 to 12", remove entry and add connections to users.json
//	else if invites does not contain "from 12 to 13", add it
func AddPoolMember(w http.ResponseWriter, r *http.Request) {
	var user Id
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("[X] Error AddPoolMembers json.Decode: ", err)
		http.Error(w, "Failed to add new member", http.StatusInternalServerError)
		return
	}

	str_target_id := r.PathValue("id")
	target_id, _ := strconv.Atoi(url.PathEscape(str_target_id))

	err = utils.AddInvite(user.User_Id, target_id)
	if err != nil {
		log.Printf("[X] Error AddPoolMembers utils.AddInvite usr%d -> usr%d: %s\n", user.User_Id, target_id, err)
		http.Error(w, "Failed to add new member", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Created new invite\n")
}

func DeletePoolMember(w http.ResponseWriter, r *http.Request) {
	var user Id
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("[X] Error DeletePoolMembers json.Decode: ", err)
		http.Error(w, "Failed to delete pool member", http.StatusInternalServerError)
		return
	}

	str_target_id := r.PathValue("id")
	target_id, _ := strconv.Atoi(url.PathEscape(str_target_id))

	err = utils.RemoveConnection(user.User_Id, target_id)
	if err != nil {
		log.Printf("[X] Error DeletePoolMembers utils.RemoveConnection usr%d -> usr%d: %s\n", user.User_Id, target_id, err)
		http.Error(w, "Failed to add new member", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Member removed\n")
}
