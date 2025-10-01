package v1

import (
	"fmt"
	"log"
	"main/utils"
	"net/http"
)

//	http.HandleFunc("GET /v1/getAllActiveUsers", nil)
//	http.HandleFunc("GET /v1/getServerStatus", nil)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	active := r.URL.Query().Get("active")
	if active == "true" {
		// iterate over wg0 interface or WGDashboard call ? Like GetAllUsers could be heavy
		fmt.Fprintf(w, "[admin] get all active users\n")
	} else {
		// iterate over all eth1.X subinterfaces and get numbers
		// could be heavy + deadlock on /etc/network/interface file...
		// maybe run a script every X that reads file and store in json ?
		stdout, stderr := utils.Exec("../scripts/createUser.sh")
		log.Println("stdout: ", stdout)
		log.Println("stderr: ", stderr)

		fmt.Fprintf(w, "[admin] get all users\n")
	}
}

func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[admin] server status\n")
}
