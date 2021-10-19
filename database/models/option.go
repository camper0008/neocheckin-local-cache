package models

type Option struct {
	DatabaseModel
	WrapperId int
	Name      string
	Available bool
}
