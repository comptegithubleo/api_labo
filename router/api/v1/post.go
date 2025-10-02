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

func ClearUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "clear user %s\n", id)
}

// if 12 adds 13 :
//
//	if invites contains "from 13 to 12", remove entry and add connections to users.json
//	else if invites does not contain "from 12 to 13", add it
func AddPoolMember(w http.ResponseWriter, r *http.Request) {
	var invites, err = utils.GetJSONInvites()
	if err != nil {
		log.Println("[X] Error AddPoolMembers utils.GetJSONInvites: ", err)
		http.Error(w, "Failed to add new member", http.StatusInternalServerError)
		return
	}

	var user Id
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("[X] Error AddPoolMembers json.Decode: ", err)
		http.Error(w, "Failed to add new member", http.StatusInternalServerError)
		return
	}

	//str_target_id := r.PathValue("id")
	//target_id, _ := strconv.Atoi(url.PathEscape(str_target_id))

	log.Println(invites)
	err = utils.AddInvite(10, 13)
	if err != nil {
		log.Println("[X] Error AddPoolMembers utils.AddInvite: ", err)
		http.Error(w, "Failed to add new member", http.StatusInternalServerError)
		return
	}

	/* for i := 0; i < len(invites); i++ {
		if invites[i].To == user.User_Id && invites[i].From == target_id {
			//remove from invite, add connection to users.json
			// https://go.dev/wiki/SliceTricks#delete-without-preserving-order because faster
			invites[i] = invites[len(invites)-1]
			invites = invites[:len(invites)-1]

			err := utils.AddConnection(user.User_Id, target_id)
			if err != nil {
				log.Println("[X] Error AddPoolMembers utils.AddInvite: ", err)
				http.Error(w, "Failed to add new member", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Added user to pool\n")
			return
		}

		if invites[i].From == user.User_Id && invites[i].To == target_id {
			http.Error(w, "Invite already exists", http.StatusInternalServerError)
			return
		}
	}

	invites = append(invites, utils.PendingInvite{
		From: user.User_Id,
		To:   target_id,
	})

	err = utils.WriteJSONInvites(invites)
	if err != nil {
		log.Println("[X] Error AddPoolMembers utils.WriteJSONInvites: ", err)
		http.Error(w, "Failed to add new member", http.StatusInternalServerError)
		return
	}
	*/
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
	err = utils.RemoveConnection(target_id, user.User_Id)
}
