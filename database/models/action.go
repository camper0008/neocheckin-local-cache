package models

import (
	"neocheckin_cache/shared"
	"time"
)

type Action struct {
	DatabaseModel
	Timestamp time.Time
	Option    shared.WrapperEnum
	Rfid      string
}
