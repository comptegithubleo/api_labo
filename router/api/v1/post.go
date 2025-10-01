package v1

import (
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create user\n")
}

func ClearUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "clear user %s\n", id)
}
