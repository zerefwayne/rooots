package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id"`
	StravaId  uint64    `json:"stravaId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}
