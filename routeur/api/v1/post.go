package v1

import (
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create user\n")

	//je ne suis pas certain mais si j'ai bien compris:

}
