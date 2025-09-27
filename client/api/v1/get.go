package v1

import (
	"fmt"
	"net/http"
)

func GetPoolMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pool members\n")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	active := r.URL.Query().Get("active")
	if active == "true" {
		// iterate over wg0 interface or WGDashboard call ? Like GetAllUsers could be heavy
		fmt.Fprintf(w, "[client] get all active users\n")
	} else {
		// iterate over all eth1.X subinterfaces and get numbers
		// could be heavy + deadlock on /etc/network/interface file...
		// maybe run a script every X that reads file and store in json ?
		fmt.Fprintf(w, "[client] get all users\n")
	}
}

func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[client] server status\n")
}
