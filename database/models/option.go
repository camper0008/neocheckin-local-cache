package models

import (
	"neocheckin_cache/shared"
)

type Option struct {
	DatabaseModel
	WrapperId shared.WrapperEnum
	Name      string
	Available bool
}
