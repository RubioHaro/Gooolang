package main

import (
	"encoding/json"
	"net/http"
)

// Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

// User ...
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// MetaData ...
type MetaData interface {
}

func (u *User) toJSON() ([]byte, error) {
	return json.Marshal(u)
}
