package model

type Registration struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"userId"`
	EventID int64 `json:"eventId"`
}
