package library

import "time"

type Order struct {
	Id           uint16
	UserId       uint16
	Book         *Book
	GiveBackDate time.Time
}
