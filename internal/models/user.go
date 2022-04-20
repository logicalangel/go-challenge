package models

import "time"

type User struct {
	UserName  string
	Segment   string
	CreatedAt time.Time
}
