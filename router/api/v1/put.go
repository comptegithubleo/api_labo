package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"main/utils"
	"net/http"
)

func ClearUser(w http.ResponseWriter, r *http.Request) {
	var user Id
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("[X] Error DeletePoolMembers json.Decode: ", err)
		http.Error(w, "Failed to delete pool member", http.StatusInternalServerError)
		return
	}

	err = utils.RemoveConnections(user.User_Id)
	if err != nil {
		log.Printf("[X] Error deleting all connections of usr%d 🎭🚀 : %s\n", user.User_Id, err)
		http.Error(w, "Failed to clear user info", http.StatusInternalServerError)
		return
	}

	log.Printf("Cleared all connections for usr%d 🛟🚨\n", user.User_Id)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User cleared\n")
}
