package transactiontemplates

import (
	"fmt"

	"github.com/Yifeeeeei/EraOfArcaneBackend/arch"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/consts"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"
)

type DealDamageTransaction struct {
	Executor   *model.Executor
	Id         model.Id
	Host       model.Id
	Target     model.Id
	States     []string
	Values     map[string]any
	InitialDmg int
	Board      *arch.Board
}

func NewDealDamageTransaction(board *arch.Board, host model.Id, target model.Id, Dmg int) *DealDamageTransaction {
	return &DealDamageTransaction{
		Id:     board.IdGenerator.GenerateId(),
		Board:  board,
		Host:   host,
		Target: target,
		States: []string{
			consts.STATE_TRANSACTION,
			consts.STATE_TRANSACTION_TYPE,
		},
		Values: map[string]any{
			consts.KEY_TRANSACTION_TYPE: consts.VALUE_TRANSACTION_DEAL_DAMAGE,
		},
		InitialDmg: Dmg,
	}
}

func (t *DealDamageTransaction) Execute(modifiers map[string]any) error {
	dmgAmount := t.InitialDmg
	if val, ok := modifiers[MODIFIER_DAMAGE_ADD_ON]; ok {
		dmgAmount += val.(int)
	}

	_, err := t.Board.GetCardById(t.Target)
	if err != nil {
		return err
	}

	t.Executor.AddTransaction(NewTakeDamageTransaction(t.Board, t.Host, t.Target, dmgAmount))
	return nil

}

func (t *DealDamageTransaction) GetStates() []string {
	return t.States
}

func (t *DealDamageTransaction) GetValues() map[string]any {
	return t.Values
}

func (t *DealDamageTransaction) GetId() model.Id {
	return t.Id
}

func (t *DealDamageTransaction) GetHost() model.Id {
	return t.Host
}

type TakeDamageTransaction struct {
	Executor   *model.Executor
	Id         model.Id
	Host       model.Id
	States     []string
	Values     map[string]any
	InitialDmg int
	Board      *arch.Board
	Source     model.Id
}

func NewTakeDamageTransaction(board *arch.Board, damageSource model.Id, host model.Id, Dmg int) *TakeDamageTransaction {
	return &TakeDamageTransaction{
		Id:    board.IdGenerator.GenerateId(),
		Board: board,
		Host:  host,
		States: []string{
			consts.STATE_TRANSACTION,
			consts.STATE_TRANSACTION_TYPE,
		},
		Values: map[string]any{
			consts.KEY_TRANSACTION_TYPE: consts.VALUE_TRANSACTION_TAKE_DAMAGE,
		},
		Source:     damageSource,
		InitialDmg: Dmg,
	}
}

func (t *TakeDamageTransaction) Execute(modifiers map[string]any) error {
	dmgAmount := t.InitialDmg
	if val, ok := modifiers[MODIFIER_DAMAGE_ADD_ON]; ok {
		dmgAmount += val.(int)
	}

	cardFom, err := t.Board.GetCardById(t.Host)
	if err != nil {
		return err
	}
	// if it is a companion or charactor
	if !(cardFom.IsCompanion() || cardFom.IsCharacter()) {
		return fmt.Errorf("only companion and character can take damage")
	}

	cardFom.Life -= dmgAmount

	if cardFom.Life <= 0 {
		t.Executor.AddTransaction(NewDieTransaction(t.Board, t.Host))
	}
	return nil

}

func (t *TakeDamageTransaction) GetStates() []string {
	return t.States
}

func (t *TakeDamageTransaction) GetValues() map[string]any {
	return t.Values
}

func (t *TakeDamageTransaction) GetId() model.Id {
	return t.Id
}

func (t *TakeDamageTransaction) GetHost() model.Id {
	return t.Host
}

//

type DieTransaction struct {
	Executor *model.Executor
	Id       model.Id
	States   []string
	Values   map[string]any
	Board    *arch.Board
	Host     model.Id
}

func NewDieTransaction(board *arch.Board, host model.Id) *DieTransaction {

	return &DieTransaction{
		Id:    board.IdGenerator.GenerateId(),
		Board: board,
		Host:  host,
		States: []string{
			consts.STATE_TRANSACTION,
			consts.STATE_TRANSACTION_TYPE,
		},
		Values: map[string]any{
			consts.KEY_TRANSACTION_TYPE: consts.VALUE_TRANSACTION_DIE,
		},
	}
}

func (t *DieTransaction) Execute(modifiers map[string]any) error {
	// put the card into graveyard
	card, err := t.Board.GetCardById(t.Host)
	if err != nil {
		return err
	}

	if _, ok := card.Values[consts.KEY_LOCATION]; !ok {
		return fmt.Errorf("card %v does not have location", t.Host)
	}

	card.Values[consts.KEY_LOCATION] = consts.VALUE_LOCATION_GRAVEYARD
	return nil
}

func (t *DieTransaction) GetStates() []string {
	return t.States
}

func (t *DieTransaction) GetValues() map[string]any {
	return t.Values
}

func (t *DieTransaction) GetId() model.Id {
	return t.Id
}

func (t *DieTransaction) GetHost() model.Id {
	return t.Host
}
