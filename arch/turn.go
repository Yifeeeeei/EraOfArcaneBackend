package arch

import (
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/consts"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"
)

// a turn is just a marker for the current turn, but it is also an instance

// states:
// - turn
// values
// - turnNumber: int

type Turn struct {
	Id         model.Id
	TurnNumber int
}

func NewTurn() *Turn {
	return &Turn{Id: model.IdGeneratorInstance.GenerateId()}
}

// implement instance interface
func (t *Turn) GetId() model.Id {
	return t.Id
}

func (t *Turn) GetStates() []string {
	return []string{consts.STATE_TURN}
}

func (t *Turn) GetValues() map[string]any {
	return map[string]any{consts.KEY_TURN_NUMBER: t.TurnNumber}
}
