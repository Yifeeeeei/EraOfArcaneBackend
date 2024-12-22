package model

type Instance interface {
	GetId() Id
	GetStates() []string
	GetValues() map[string]any
}
