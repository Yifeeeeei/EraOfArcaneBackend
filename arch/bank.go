package arch

import "github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"

// bank holds the elements of ONE player
type Bank struct {
	Id       model.Id
	Elements *Elements
}

// Bank is an instance
// states:
// - bank
// values:
// - it does not provide values

func NewBank() *Bank {
	return &Bank{
		Id:       model.IdGeneratorInstance.GenerateId(),
		Elements: NewElements(),
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
