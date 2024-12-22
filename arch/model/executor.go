package model

type Executor struct {
	TransactionQueue []Transaction
	EffectsBefore    []Effect
	EffectsAfter     []Effect
	currentIndex     int
}

func NewExecutor() *Executor {
	return &Executor{
		TransactionQueue: make([]Transaction, 0),
		currentIndex:     0,
	}
}

func (e *Executor) AddTransaction(t Transaction) {
	e.TransactionQueue = append(e.TransactionQueue, t)
}

func (e *Executor) AddEffectBefore(effect Effect) {
	e.EffectsBefore = append(e.EffectsBefore, effect)
}

func (e *Executor) AddEffectAfter(effect Effect) {
	e.EffectsAfter = append(e.EffectsAfter, effect)
}

func (e *Executor) ExecuteAll() error {
	for e.currentIndex < len(e.TransactionQueue) {
		t := e.TransactionQueue[e.currentIndex]
		params := map[string]any{}
		for _, effect := range e.EffectsBefore {
			var err error
			params, err = effect.Modify(t, params)
			if err != nil {
				return err
			}
		}

		err := t.Execute(params)
		if err != nil {
			return err
		}

		for _, effect := range e.EffectsAfter {
			var err error
			params, err = effect.Modify(t, params)
			if err != nil {
				return err
			}
		}

		e.currentIndex++
	}
	return nil
}
