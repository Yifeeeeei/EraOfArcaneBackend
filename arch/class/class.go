package class

// 所有的分类

type Class int

// 技能分类: 创造，驱动，灵媒，聚能，神秘，幻变
const (
	Create Class = iota
	Drive
	Spirit
	Aggregate
	Mystery
	Metamorphosis
)

// 伙伴分类：巫师，人类，龙，造物，野兽，异兽，恶魔，精灵
const (
	Wizard Class = iota
	Human
	Dragon
	Creature
	Beast
	Monster
	Demon
	Elf
)

// 道具分类 装备，消耗品
const (
	Equipment Class = iota
	Consumable
	Potion
	Scroll
	SpellScroll
)

// 稀有度
const (
	Legendary Class = iota
	Spawn
)
