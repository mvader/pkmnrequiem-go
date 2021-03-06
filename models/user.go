package models

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
)

type User struct {
	ID      string   `json:"id"`
	Badges  int      `json:"badges"`
	Battles []string `json:"battles"`
	Email   string   `json:"email"`
	Money   int      `json:"money"`
	// Party []Pokemon `json:"party"`
	Password string `json:"password"`
	Username string `json:"username"`
	Verified bool   `json:"verified"`
}

type UserStore struct {
	*mgo.Database
}

func (s UserStore) ByToken(token string) (*User, error) {
	return nil, nil
}

func NewUser(email, password, username string) *User {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user := &User{
		ID: uuid.NewV4().String(),
		// Badges: 0,
		// Battles: []string{},
		Email: email,
		Money: 3000,
		// Party: []Pokemon{},
		Password: string(encryptedPassword),
		Username: username,
		// Verified: false,
	}
	return user
}
