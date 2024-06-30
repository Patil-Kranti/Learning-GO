package main

import (
	"math/rand"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type User struct {
	Id                int       `json:"id"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"-"`
	Token             string    `json:"token"`
	AccountId         int       `json:"accountId"`
	CreatedAt         time.Time `json:"createdAt"`
}
type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}
type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000000)),
		CreatedAt: time.Now().UTC(),
	}
}
func NewUser(email string, password string, token string, accountId int) *User {
	return &User{
		Email:             email,
		EncryptedPassword: password,
		Token:             token,
		AccountId:         accountId,
		CreatedAt:         time.Now().UTC(),
	}
}
