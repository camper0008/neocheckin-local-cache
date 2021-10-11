package database_models

import "time"

type Action struct {
	Time       time.Time
	OptionId   int
	CheckingIn bool
	Rfid       string
}
