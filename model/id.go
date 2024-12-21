package model

type Id struct {
	id    int
	valid bool
}

type IdGenerator struct {
	currentId int
}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{
		currentId: 0,
	}
}

func (ig *IdGenerator) GenerateId() Id {
	ig.currentId++
	return Id{ig.currentId, true}
}

var IdGeneratorInstance = NewIdGenerator()
