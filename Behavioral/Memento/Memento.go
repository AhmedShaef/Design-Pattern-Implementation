package memento

type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

type originator struct {
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{state: o.state}
}

func (o *originator) restoreMemento(m *memento) {
	o.state = m.getSavedState()
}
func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}

type caretaker struct {
	mementoSlice []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoSlice = append(c.mementoSlice, m)
}

func (c *caretaker) getMemento(idx int) *memento {
	return c.mementoSlice[idx]
}
