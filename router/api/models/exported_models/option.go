package exported_models

import (
	"neocheckin_cache/shared"
)

type Option struct {
	Id        shared.WrapperEnum `json:"id"`
	Name      string             `json:"name"`
	Available bool               `json:"available"`
}
