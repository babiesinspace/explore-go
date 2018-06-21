package api

import "time"

type Match struct {
	ID           int       `json:"id"`
	SenderID     int       `json:"senderId"`
	ReceiverID   int       `json:"receiverId"`
	Comment      string    `json:"comment"`
	Sent         time.Time `json:"sent"`
	Reciprocated time.Time `json:"reciprocated"`
}
