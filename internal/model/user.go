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

func (u UserId) String() string {
	return uuid.UUID(u).String()
}

func UserIdFromString(s string) UserId {
	return UserId(uuid.MustParse(s))
}
