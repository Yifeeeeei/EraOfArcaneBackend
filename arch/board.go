package arch

import (
	"fmt"

	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/consts"
	"github.com/Yifeeeeei/EraOfArcaneBackend/arch/model"
)

// Board is a collection of all instances, marking everything in this game.
// If you want to find something, you should ask board with the corresponding id.

type Board struct {
	IdGenerator model.IdGenerator

	AllInstances map[model.Id]model.Instance
	Player0Id    model.Id
	Player1Id    model.Id
	// a card has to be in one of the following places
	Player0CardIds []model.Id
	Player1CardIds []model.Id
	NeutralCardIds []model.Id
	// the turn
	Player0BankId model.Id
	Player1BankId model.Id

	TurnId model.Id
}

func NewBoard(player0Bank *Bank, player1Bank *Bank, turn *Turn) *Board {
	return &Board{
		AllInstances:   make(map[model.Id]model.Instance),
		Player0CardIds: make([]model.Id, 0),
		Player1CardIds: make([]model.Id, 0),
		NeutralCardIds: make([]model.Id, 0),
		Player0BankId:  player0Bank.GetId(),
		Player1BankId:  player1Bank.GetId(),
		TurnId:         turn.GetId(),
	}
}

func (b *Board) AddCard(card *Card) error {
	// since all card are instances, add it to the board
	b.AllInstances[card.GetId()] = card
	// add the card to the corresponding place
	switch card.GetValues()[consts.STATE_OWNER] {
	case consts.VALUE_OWNER_PLAYER0:
		b.Player0CardIds = append(b.Player0CardIds, card.GetId())
	case consts.VALUE_OWNER_PLAYER1:
		b.Player1CardIds = append(b.Player1CardIds, card.GetId())
	case consts.VALUE_OWNER_NEUTRAL:
		b.NeutralCardIds = append(b.NeutralCardIds, card.GetId())
	default:
		return fmt.Errorf("unknown owner: %s", card.GetValues()[consts.STATE_OWNER])
	}
	return nil
}

func (b *Board) RemoveCardById(cardId model.Id) error {
	// remove the card from the board
	delete(b.AllInstances, cardId)
	// remove the card from the corresponding place
	for i, id := range b.Player0CardIds {
		if id == cardId {
			b.Player0CardIds = append(b.Player0CardIds[:i], b.Player0CardIds[i+1:]...)
			return nil
		}
	}
	for i, id := range b.Player1CardIds {
		if id == cardId {
			b.Player1CardIds = append(b.Player1CardIds[:i], b.Player1CardIds[i+1:]...)
			return nil
		}
	}
	for i, id := range b.NeutralCardIds {
		if id == cardId {
			b.NeutralCardIds = append(b.NeutralCardIds[:i], b.NeutralCardIds[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("card not found: %s", cardId.String())
}

func (b *Board) GetCardById(cardId model.Id) (*Card, error) {
	card, ok := b.AllInstances[cardId].(*Card)
	if !ok {
		return nil, fmt.Errorf("instance not found: %s", cardId.String())
	}
	return card, nil
}
