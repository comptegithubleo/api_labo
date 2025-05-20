package api_v1

import (
	"encoding/json"
	"log"
	"main/data"
	"main/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		log.Printf("Failed to generate UUID: %v", err)
		return
	}
	user.UUID = uuid.String()

	// validate inputs
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(user)
	if err != nil {
		log.Printf("Validation failed: %v", err)
		http.Error(w, "Input validation failed. Verify username, password and mail requirements", http.StatusBadRequest)
		return
	}

	// create user in db
	err = data.CreateUser(user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, "Server failed to create user (duplicated email or internal error)", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully with UUID: " + user.UUID))
}
