package model

// a transaction is a single operation that can be executed on an instance
type Transaction interface {
	Execute(map[string]any) error
	GetTypes() []string
	GetParams() map[string]any
	GetId() Id
	GetFrom() Id
	GetTo() Id
}
