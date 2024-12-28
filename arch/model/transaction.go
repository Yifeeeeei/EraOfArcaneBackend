package model

// a transaction is a single operation that can be executed on an instance
type Transaction interface {
	Execute(map[string]any) error
	GetStates() []string
	GetValues() map[string]any
	GetId() Id
	GetHost() Id
}
