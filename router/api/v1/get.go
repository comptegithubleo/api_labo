package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//	http.HandleFunc("GET /v1/getAllActiveUsers", nil)
//	http.HandleFunc("GET /v1/getServerStatus", nil)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data/users.json")
	if err != nil {
		log.Println("[X] Error GetUsers: ", err)
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetPendingConnections(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data/pending.json")
	if err != nil {
		log.Println("[X] Error GetUsers: ", err)
		http.Error(w, "Failed to retrieve pending connections", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[admin] server status\n")
}
