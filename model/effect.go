package model

type Effect interface {
	GetId() Id
	// an affect can modity a transaction
	Modify(Transaction, map[string]any) (map[string]any, error)
}
