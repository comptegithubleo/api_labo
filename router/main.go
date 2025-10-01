package main

//copier coller de main.go
import (
	"fmt"
	"log"
	api "main/api/v1"
	"net/http"
)

func main() {

	http.HandleFunc("GET /v1/users", api.GetUsers)
	http.HandleFunc("GET /v1/invites", api.GetPendingConnections)
	http.HandleFunc("POST /v1/users", api.CreateUser)
	http.HandleFunc("PUT /v1/users/{id}", api.ClearUser) //reset user info & pool
	http.HandleFunc("DELETE /v1/users/{id}", api.DeleteUser)

	http.HandleFunc("POST /v1/pool/members/{id}", api.AddPoolMember) //add member to pool
	http.HandleFunc("DELETE /v1/pool/members/{id}", api.DeletePoolMember)


	http.HandleFunc("GET /v1/status", api.Status) //server status

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Landing page admin api\n")
	})

	log.Println("Starting router api server on 3032")
	http.ListenAndServe(":3032", nil)
}
