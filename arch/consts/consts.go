package consts

// -- vocabulary --
// 伙伴 companion
// 技能 ability
// 物品 item
// 人物 character
// 生命 life
// 攻击 attack
// 花费 cost
// 入场 enter
// 使用 use
// 负载 gain
// 威力 power
// 属性 elem
// 法术 spell
// 咒术 curse

// states
// here list all the states and its corresponding values

// turn.go
const (
	STATE_TURN      = "turn_turn"
	KEY_TURN_NUMBER = "turn_turnNumber" // int
)

// card

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
)

// transactions
const (
	STATE_TRANSACTION = "transaction"

	STATE_TRANSACTION_TYPE        = "transaction_type"
	KEY_TRANSACTION_TYPE          = "transaction_type" // string, half-defined
	VALUE_TRANSACTION_DEAL_DAMAGE = "transaction_deal_damage"
	VALUE_TRANSACTION_TAKE_DAMAGE = "transaction_take_damage"
	VALUE_TRANSACTION_DIE         = "transaction_die"
)

// dual
const (
	STATE_DUAL = "dual"
)
