package exported_models

type OptionAvailable int

const (
	NOT_AVAILABLE = 0
	AVAILABLE     = 1
	PRIORITY      = 2
	INVALID       = 3
)

type Option struct {
	Id          int             `json:"id"`
	Name        string          `json:"name"`
	DisplayName string          `json:"displayName"`
	Category    string          `json:"category"`
	Available   OptionAvailable `json:"available"`
}
