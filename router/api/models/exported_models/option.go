package exported_models

type OptionAvailable int

const (
	NOT_AVAILABLE = 0
	AVAILABLE     = 1
	PRIORITY      = 2
)

type Option struct {
	Id        int             `json:"id"`
	Name      string          `json:"name"`
	Available OptionAvailable `json:"available"`
}
