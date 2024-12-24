package arch

import (
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/element"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"
)

// bank holds the elements of ONE player
type Bank struct {
	Id       model.Id
	Elements *element.Elements
}

// Bank is an instance
// states:
// - bank
// values:
// - it does not provide values

func NewBank(board *Board) *Bank {
	return &Bank{
		Id:       board.IdGenerator.GenerateId(),
		Elements: element.NewElements(),
	}
}

func (b *Bank) GetId() model.Id {
	return b.Id
}

func (b *Bank) GetStates() []string {
	return []string{"bank"}
}

func (b *Bank) GetValues() map[string]any {
	return map[string]any{}
}
