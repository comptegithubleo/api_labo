package main

import (
	"fmt"
	"log"
	api "main/api/v1"
	"net/http"
)

func main() {

	http.HandleFunc("GET /v1/users", api.GetUsers)
	http.HandleFunc("GET /v1/invites", api.GetInvites)
	http.HandleFunc("PUT /v1/users/me", api.ClearUser)

	// Deprecated because available in GetUsers data ?
	//http.HandleFunc("GET /v1/pool/members", api.GetPoolMembers)

	http.HandleFunc("POST /v1/pool/add/{id}", api.AddPoolMember) //add member to pool
	http.HandleFunc("POST /v1/pool/delete/{id}", api.DeletePoolMember)

	http.HandleFunc("GET /v1/status", api.Status) //server status

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Landing page client api\n")
	})

	log.Println("Starting client api server on 3030")
	http.ListenAndServe(":3030", nil)
}
