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

type Foam struct {
	Coffe Coffe
}

func (f *Foam) Cost() float64 {

	return f.Coffe.Cost() + 1.0
}
func (f *Foam) Description() string {
	return f.Coffe.Description() + ". Foam"
}

func (m *Milk) AddFoam() {
	m.Coffe = &Foam{Coffe: m.Coffe}
}
