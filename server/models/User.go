package models

type User struct {
	ID        int64  `json:"id"`
	StravaId  int64  `json:"stravaId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
