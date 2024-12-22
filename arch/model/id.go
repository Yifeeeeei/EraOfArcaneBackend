package model

// if you want to use id, you should always call IdGeneratorInstance.GenerateId()
var IdGeneratorInstance = NewIdGenerator()

type Id struct {
	id    int
	valid bool
}

func (i Id) IsValid() bool {
	return i.valid
}

func (i Id) SameAs(other Id) bool {
	if i.valid && other.valid {
		return i.id == other.id
	} else {
		return false
	}
}

func (i Id) String() string {
	if i.valid {
		return string(i.id)
	} else {
		return "invalid_id"
	}
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
