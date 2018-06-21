package api

import "time"

type Like struct {
	ID         int       `json:"id"`
	SenderID   int       `json:"senderId"`
	ReceiverID int       `json:"receiverId"`
	Comment    string    `json:"comment"`
	Created    time.Time `json:"created"`
}
