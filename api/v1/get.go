package api_v1

import (
	"encoding/json"
	"main/data"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data.GetUsers)
}
