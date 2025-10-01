package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create user\n")

	password := os.Getenv("PROXMOX_PASSWORD")

	log.Println(password)

	/* data := bytes.NewBufferString(`{"hello":"world","answer":42}`)
	req, _ := http.NewRequest("PUT", "http://10.0.0.1:8006", data)
	req.Header.Set("Content-Type", "application/json")
	command, _ := http2curl.GetCurlCommand(req)
	fmt.Println(command) */
}

func ClearUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "clear user %s\n", id)
}
