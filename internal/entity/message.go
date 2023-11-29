package entity

import "time"

type Message struct {
	Id          string
	Name        string
	ContactInfo string
	Body        string
	Date        time.Duration
}
