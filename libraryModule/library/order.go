package library

import "time"

type Order struct {
	id           uint16
	UserId       uint16
	Book         Book
	giveBackDate time.Time
}
