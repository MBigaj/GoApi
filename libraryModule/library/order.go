package library

import "time"

type Order struct {
	Id           uint16
	User         *User
	Book         *Book
	GiveBackDate time.Time
	Completed    bool
}
