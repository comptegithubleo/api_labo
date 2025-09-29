package v1

import (
	"fmt"
	"net/http"
)

func ClearUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "clear user %s\n", id)
}
