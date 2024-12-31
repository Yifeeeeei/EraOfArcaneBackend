package arch

import "github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"

// 发起一次法术对抗

// Dual is an instance
type Dual struct {
	Id              model.Id
	States          []string
	Values          map[string]any
	MainSpellId     model.Id
	EnhanceSpellIds []model.Id
	Modifier        map[string]any
}

// implement instance interface
func (d *Dual) GetId() model.Id {
	return d.Id
}

func (d *Dual) GetStates() []string {
	return d.States
}

func (d *Dual) GetValues() map[string]any {
	return d.Values
}
