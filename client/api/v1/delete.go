package v1

import (
	"fmt"
	"net/http"
)

func DeletePoolMember(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "delete pool member %s\n", id)
}
