package arch

import (
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/class"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/consts"
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

// companion card is a type of card, it should have life, attack
type CompanionCard struct {
	Card
	Life   int
	Attack int
	Gain   element.Elements
	Class  class.Class
}

// ability card
type AbilityCard struct {
	Card
	UseCost element.Elements
}

type SpellCard struct {
	AbilityCard
	Attack int
	Power  int
}

type CurseCard struct {
	AbilityCard
}

// item card, this is the complicated one
type ItemCard struct {
	Card
}

type EquipmentCard struct {
	ItemCard
	Attack int
	Gain   element.Elements
}

func NewEquipmentCard(board *Board, elem element.Elem, enterCost element.Elements, rarity rarity.Rarity, attack int, gain element.Elements, value_owner string, value_location string) *EquipmentCard {
	return &EquipmentCard{
		ItemCard: ItemCard{
			Card: Card{
				Board:  board,
				States: []string{consts.STATE_CARD, consts.KEY_OWNER, consts.STATE_LOCATION, consts.STATE_TYPE},
				Values: map[string]any{
					consts.KEY_OWNER:    value_owner,
					consts.KEY_LOCATION: value_location,
					consts.KEY_TYPE:     consts.VALUE_TYPE_ITEM,
				},
				Id:        board.IdGenerator.GenerateId(),
				EnterCost: enterCost,
				Elem:      elem,
				Classes:   []class.Class{class.Equipment},
				Rarity:    rarity,
			},
		},
	}
}

type ConsumableCard struct {
	ItemCard
}

type SpellScrollCard struct {
	ConsumableCard
	Attack int
	Power  int
}

func NewSpellScrollCard(board *Board, elem element.Elem, enterCost element.Elements, rarity rarity.Rarity, value_owner string, value_location string) *SpellScrollCard {
	return &SpellScrollCard{
		ConsumableCard: ConsumableCard{
			ItemCard: ItemCard{
				Card: Card{
					Board:  board,
					States: []string{consts.STATE_CARD, consts.KEY_OWNER, consts.STATE_LOCATION, consts.STATE_TYPE},
					Values: map[string]any{
						consts.KEY_OWNER:    value_owner,
						consts.KEY_LOCATION: value_location,
						consts.KEY_TYPE:     consts.VALUE_TYPE_ITEM,
					},
					Id:        board.IdGenerator.GenerateId(),
					EnterCost: enterCost,
					Elem:      elem,
					Classes:   []class.Class{class.Consumable, class.SpellScroll},
					Rarity:    rarity,
				},
			},
		},
	}
}

type OtherConsumableCard struct {
	ConsumableCard
}

// character card
type CharacterCard struct {
	Card
	Life   int
	Attack int
	Gain   element.Elements
}
