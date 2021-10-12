package exported_models

import (
	"neocheckin_cache/shared"
)

type Option struct {
	Id        shared.WrapperEnum
	Name      string
	Available bool
}
