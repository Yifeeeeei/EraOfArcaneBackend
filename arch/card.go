package arch

import (
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/class"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/element"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/rarity"
)

// this is the biggest part of arch
// 1. every card is an instance
// 2. we need to take account of every possible state:
//    - whose card it is? (owner: player0, player1, neutral)
//    - where is it? (location: deck, hand, battlefield, graveyard)
//    - what type is it? (type: companion, ability, item, character)
//    - register all these in consts

// relations
// card --> ability
//      |   |-> spell
//      |   |-> curse
//      |-> companion
//      |-> item
//	    |   |-> equipment
//	    |   |-> consumable
//      |-> character

type Card struct {
	Board     *Board // have the board pointer to access other stuffs
	Id        model.Id
	States    []string
	Values    map[string]any
	EnterCost element.Elements
	Elem      element.Elem
	Classes   []class.Class
	Rarity    rarity.Rarity
	// ability
	Duration int
	Power    int
	// companion
	Attack int
	Life   int
}

func (c *Card) GetId() model.Id {
	return c.Id
}

func (c *Card) GetStates() []string {
	return c.States
}

func (c *Card) GetValues() map[string]any {
	return c.Values
}

func (c *Card) GetEnterCost() element.Elements {
	return c.EnterCost
}
