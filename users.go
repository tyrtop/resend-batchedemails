package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type User struct {
	Email string `json:"Email"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func loadUsers(path string) ([]User, error) {
	var users []User

	data, err := os.ReadFile(path)
	if err != nil{
		return nil, fmt.Errorf("loading file from %s: %w", path, err)
		
	}
	err =json.Unmarshal(data, &users)
	if err != nil {
		return nil,fmt.Errorf("loading users from %s: %w", path, err)
	}
	return users, nil
}
