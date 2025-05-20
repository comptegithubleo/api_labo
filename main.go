package main

import (
	api "main/api/v1"
	"main/data"
	"net/http"

	"github.com/gorilla/mux"
)

/* func CheckLoginMiddleware(db *sql.DB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uuid := r.Header.Get("UUID")
			if uuid == "" {
				http.Error(w, "UUID header missing", http.StatusUnauthorized)
				return
			}

			userExists, err := checkUserExists(db, uuid)
			if err != nil {
				http.Error(w, "Database error", http.StatusInternalServerError)
				log.Printf("Database error: %v", err)
				return
			}

			if !userExists {
				http.Error(w, "Not logged in", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getUserInfoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuidParam := r.Header.Get("UUID")
		name, email, err := getUserInfo(db, uuidParam)
		if err != nil {
			http.Error(w, "Failed to retrieve user information", http.StatusInternalServerError)
			log.Printf("Failed to retrieve user information: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"name": "%s", "email": "%s"}`, name, email)))
	}
} */

func main() {
	r := mux.NewRouter()
	data.Connect()

	r.HandleFunc("/users", api.GetUsers).Methods("GET") // Public route
	r.HandleFunc("/user", api.CreateUser).Methods("POST")  // Public route

	http.ListenAndServe(":8080", r)
}
