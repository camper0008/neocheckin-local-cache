package exported_models

import (
	"neocheckin_cache/shared"
	"time"
)

type Action struct {
	Timestamp time.Time
	Option    shared.WrapperEnum
	Rfid      string
}
