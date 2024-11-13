package coffe

type Milk struct {
	Coffe Coffe
}

func (m *Milk) Cost() float64 {
	return m.Coffe.Cost() + 0.5
}

func (m *Milk) Description() string {
	return m.Coffe.Description() + ". Milk"
}
