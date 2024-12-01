package model

import "github.com/google/uuid"

type User struct {
	Username   string
	FirstName  string
	MiddleName string
	LastName   string
	Snils      string
	Email      string
	Password   string
	IsAdmin    bool
}

type UserId uuid.UUID
