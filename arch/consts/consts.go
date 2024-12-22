package consts

// states
// here list all the states and its corresponding values

// turn.go
const (
	STATE_TURN      = "turn_turn"
	KEY_TURN_NUMBER = "turn_turnNumber" // int
)

// card.go

const (
	STATE_CARD = "card"

	STATE_OWNER         = "card_owner"
	KEY_OWNER           = "card_owner" // string, defined
	VALUE_OWNER_PLAYER0 = "card_player0"
	VALUE_OWNER_PLAYER1 = "card_player1"
	VALUE_OWNER_NEUTRAL = "card_neutral"

	STATE_LOCATION             = "card_location"
	KEY_LOCATION               = "card_location" // string, defined
	VALUE_LOCATION_DECK        = "card_deck"
	VALUE_LOCATION_HAND        = "card_hand"
	VALUE_LOCATION_BATTLEFIELD = "card_battlefield"
	VALUE_LOCATION_GRAVEYARD   = "card_graveyard"

	STATE_TYPE           = "card_type"
	KEY_TYPE             = "card_type" // string, defined
	VALUE_TYPE_COMPANION = "card_companion"
	VALUE_TYPE_ABILITY   = "card_ability"
	VALUE_TYPE_ITEM      = "card_item"
	VALUE_TYPE_CHARACTER = "card_character"
)
