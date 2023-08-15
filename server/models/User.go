package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `json:"id,omitempty"`
	RefreshToken string    `json:"refreshToken,omitempty"`

	StravaId      uint64    `json:"stravaId,omitempty"`
	FirstName     string    `json:"firstname,omitempty"`
	LastName      string    `json:"lastname,omitempty"`
	ProfileMedium string    `json:"profileMediumUrl,omitempty"`
	Profile       string    `json:"profileUrl,omitempty"`
	City          string    `json:"city,omitempty"`
	State         string    `json:"state,omitempty"`
	Country       string    `json:"country,omitempty"`
	Sex           string    `json:"sex,omitempty"`
	Summit        bool      `json:"summit,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
}
