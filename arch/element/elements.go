package element

// elements is not an instance

type Elem int

const (
	None Elem = iota
	Fire
	Water
	Earth
	Air
	Light
	Dark
)

type Elements struct {
	None  int
	Fire  int
	Water int
	Earth int
	Air   int
	Light int
	Dark  int
}

func NewElements() *Elements {
	return &Elements{
		None:  0,
		Fire:  0,
		Water: 0,
		Earth: 0,
		Air:   0,
		Light: 0,
		Dark:  0,
	}
}

func (e *Elements) Add(elements *Elements) {
	e.None += elements.None
	e.Fire += elements.Fire
	e.Water += elements.Water
	e.Earth += elements.Earth
	e.Air += elements.Air
	e.Light += elements.Light
	e.Dark += elements.Dark
}

func (e *Elements) Copy() *Elements {
	return &Elements{
		None:  e.None,
		Fire:  e.Fire,
		Water: e.Water,
		Earth: e.Earth,
		Air:   e.Air,
		Light: e.Light,
		Dark:  e.Dark,
	}
}
