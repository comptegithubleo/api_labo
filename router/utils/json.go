package utils

import (
	"encoding/json"
	"errors"
	"os"
)

type User struct {
	ID          int   `json:"id"`
	Connections []int `json:"connections"`
	Active      bool  `json:"active"`
}

type PendingInvite struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type PendingInvites []PendingInvite

func GetJSONInvites() (PendingInvites, error) {
	data, err := os.ReadFile("data/invites.json")
	if err != nil {
		return nil, errors.New("Opening file invites failed")
	}
	var invites PendingInvites
	json.Unmarshal(data, &invites)
	return invites, nil
}

func WriteJSONInvites(invites PendingInvites) error {
	data, err := json.Marshal(invites)
	if err != nil {
		return errors.New("Failed to Marshal JSON user file")
	}

	err = os.WriteFile("data/invites.json", data, 0644)
	if err != nil {
		return errors.New("Failed to WriteFile json invites file")
	}

	return nil
}

func GetJSONUsers() ([]User, error) {
	data, err := os.ReadFile("data/users.json")
	if err != nil {
		return nil, errors.New("Opening file users failed")
	}
	var users []User
	json.Unmarshal(data, &users)
	return users, nil
}

func WriteJSONUsers(users []User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return errors.New("Failed to Marshal JSON user file")
	}

	err = os.WriteFile("data/users.json", data, 0644)
	if err != nil {
		return errors.New("Failed to WriteFile json users file")
	}

	return nil
}
