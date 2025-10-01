package v1

import (
	"fmt"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "delete user %s\n", id)
}

func DeletePoolMember(w http.ResponseWriter, r *http.Request) {

}
