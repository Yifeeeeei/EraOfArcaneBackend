package model

import (
	"fmt"
	"testing"
)

// story:
// there are four cards
// # A has life = 2, B has (all demon damages + 1), C has 1 life and effect (when a unit dies, it gains 1 life), D has onEnter, deal 1 damage

// the game
var game = []Instance{}
var executor = NewExecutor()

type Companion interface {
	GetLife() int
}

// A
type CardA struct {
	Id   Id
	Life int
	// States []string
	// Values map[string]any
}

func NewCardA() *CardA {
	return &CardA{
		Life: 2,
	}
}

func (c *CardA) GetId() Id {
	return Id{0}
}

func (c *CardA) GetLife() int {
	return c.Life
}

func (c *CardA) GetStates() []string {
	return []string{"demon"}
}

func (c *CardA) GetValues() map[string]any {
	return map[string]any{}
}

// B's effect

type EffectB struct{}

func (e *EffectB) GetId() Id {
	return Id{1}
}

func (e *EffectB) Modify(t Transaction, modifiers map[string]any) (map[string]any, error) {
	// this effects triggers before a transaction
	transactionTypes := t.GetTypes()
	for _, transactionType := range transactionTypes {
		if transactionType == "damage" {
			_, ok := modifiers["damageAddOn"]
			if ok {
				modifiers["damageAddOn"] = modifiers["damageAddOn"].(int) + 1
			} else {
				modifiers["damageAddOn"] = 1
			}
		}
		break
	}
	return modifiers, nil
}

// C
type CardC struct {
	Id   Id
	Life int
}

func NewCardC() *CardC {
	return &CardC{
		Life: 1,
	}
}

func (c *CardC) GetId() Id {
	return Id{2}
}

func (c *CardC) GetStates() []string {
	return []string{}
}

func (c *CardC) GetValues() map[string]any {
	return map[string]any{}
}

type EffectC struct{}

func (e *EffectC) GetId() Id {
	return Id{3}
}

type AddLifeTransaction struct{}

func (t *AddLifeTransaction) Execute(modifiers map[string]any) error {
	for _, instance := range game {
		if instance.GetId().id == 2 {
			// convert it to card c
			cardC, ok := instance.(*CardC)
			if ok {
				cardC.Life++
			}

		}
	}
	return nil
}

func (t *AddLifeTransaction) GetTypes() []string {
	return []string{"addLife"}
}

func (t *AddLifeTransaction) GetId() Id {
	return Id{4}
}

func (t *AddLifeTransaction) GetFrom() Instance {
	for _, instance := range game {
		if instance.GetId().id == 2 {
			// convert it to card c
			return instance

		}
	}
	return nil
}

func (t *AddLifeTransaction) GetTo() Instance {
	for _, instance := range game {
		if instance.GetId().id == 2 {
			// convert it to card c
			return instance

		}
	}
	return nil
}

func (e *EffectC) Modify(t Transaction, modifiers map[string]any) (map[string]any, error) {
	// this effect triggers after a transaction
	// if type contains die
	transactionTypes := t.GetTypes()
	for _, transactionType := range transactionTypes {
		if transactionType == "die" {
			// add life
			executor.AddTransaction(&AddLifeTransaction{})
		}
	}
	return modifiers, nil
}

// transaction die

type DieTransaction struct {
	Id Id
}

func (t *DieTransaction) Execute(modifiers map[string]any) error {
	// remove from game
	for i, instance := range game {
		if instance.GetId().id == t.Id.id {
			game = append(game[:i], game[i+1:]...)
		}
	}
	return nil
}

func (t *DieTransaction) GetTypes() []string {
	return []string{"die"}
}

func (t *DieTransaction) GetId() Id {
	return Id{5}
}

func (t *DieTransaction) GetFrom() Instance {
	// doesn't matter here
	return nil
}

func (t *DieTransaction) GetTo() Instance {
	// doesn't matter here
	return nil
}

// D's enter transaction
type DealDamageTransaction struct {
	FromId       Id
	ToId         Id
	DamageAmount int
}

func (t *DealDamageTransaction) Execute(modifiers map[string]any) error {
	// a simplified version
	// deal damage
	dmg := t.DamageAmount
	if _, ok := modifiers["damageAddOn"]; ok {
		dmg += modifiers["damageAddOn"].(int)
	}
	for _, instance := range game {
		if instance.GetId().id == t.ToId.id {
			companion, ok := instance.(Companion)
			if ok {
				if companion.GetLife() <= dmg {
					// die
					executor.AddTransaction(&DieTransaction{Id: t.ToId})
				}
			}
		}
	}
	return nil
}

func (t *DealDamageTransaction) GetTypes() []string {
	return []string{"damage"}
}

func (t *DealDamageTransaction) GetId() Id {
	return Id{6}
}

func (t *DealDamageTransaction) GetFrom() Instance {
	// doesn't matter here
	return nil
}

func (t *DealDamageTransaction) GetTo() Instance {
	// doesn't matter here
	return nil
}

// func (c *)

func TestScene(t *testing.T) {
	fmt.Println("TestScene")

	// a is now in game
	game = append(game, NewCardA())
	// register b's effect
	executor.AddEffectBefore(&EffectB{})
	// register c's effect
	executor.AddEffectAfter(&EffectC{})
	// add c into the game
	game = append(game, NewCardC())

	// d enters, add transaction
	executor.AddTransaction(&DealDamageTransaction{
		FromId:       Id{6},
		ToId:         Id{0},
		DamageAmount: 1,
	})

	// execute all
	executor.ExecuteAll()

	for _, instance := range game {
		fmt.Println(instance.GetId())
		if instance.GetId().id == 2 {
			// print life
			cardC, ok := instance.(*CardC)
			if ok {
				fmt.Println(cardC.Life)
			}
		}
	}

}
