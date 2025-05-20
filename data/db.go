package data

import (
	"database/sql"
	"log"
	"main/models"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Connect() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/api")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
}

func GetDB() *sql.DB {
	return db
}

func checkUserExists(db *sql.DB, email string) (bool, error) {
	var userExists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&userExists)
	return userExists, err
}

func getUserInfo(db *sql.DB, uuid string) (string, string, error) {
	var name, email string
	err := db.QueryRow("SELECT name, email FROM users WHERE uuid = ?", uuid).Scan(&name, &email)
	return name, email, err
}

func GetUsers() ([]models.User, error) {
	rows, err := db.Query("SELECT uuid, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UUID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(user models.User) error {
	_, err := db.Exec("INSERT INTO users (uuid, username, email, password) VALUES (?, ?, ?, ?)", user.UUID, user.Username, user.Email, user.Password)
	return err
	/* id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("create user: %v", err)
	}
	log.Printf("User created with ID %d\n", id)
	return id, nil */
}
