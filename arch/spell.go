package arch

import "github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"

// this files defines an interface that spell cards should follow

type Spell interface {
	GetPower() int
	GetAttack() int
	GetRange() []Field
	CreateAttackTransaction() model.Transaction
	Enhance(map[string]any) (map[string]any, error)
}
