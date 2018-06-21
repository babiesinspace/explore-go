package api

import (
	"time"
)

type User struct {
	ID      int
	Name    string
	Photo   string
	Age     int
	Created time.Time
	Updated time.Time
}
