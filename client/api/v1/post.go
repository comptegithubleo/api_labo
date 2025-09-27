package v1

import (
	"fmt"
	"net/http"
)

func AddPoolMember(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "add pool member %s\n", id)
}
