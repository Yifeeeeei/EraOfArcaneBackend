package arch

import (
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/consts"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"
)

// this is the biggest part of arch
// 1. every card is an instance
// 2. we need to take account of every possible state:
//    - whose card it is? (owner: player0, player1, neutral)
//    - where is it? (location: deck, hand, battlefield, graveyard)
//    - what type is it? (type: companion, ability, item, character)
//    - register all these in consts

type Card struct {
	Board  *Board // have the board pointer to access other stuffs
	Id     model.Id
	States []string
	Values map[string]any
}

func NewCard(board *Board, cardOwner string, cardLocation string, cardType string) *Card {
	return &Card{
		Board:  board,
		Id:     model.IdGeneratorInstance.GenerateId(),
		States: []string{consts.STATE_CARD, consts.STATE_OWNER, consts.STATE_LOCATION, consts.STATE_TYPE},
		Values: map[string]any{consts.STATE_OWNER: cardOwner, consts.STATE_LOCATION: cardLocation, consts.STATE_TYPE: cardType},
	}
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
