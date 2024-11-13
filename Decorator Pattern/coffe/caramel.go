package coffe

type Caramel struct {
	Coffe Coffe
}

func (c *Caramel) Cost() float64 {
	return c.Coffe.Cost() + 1
}

func (c *Caramel) Description() string {
	return c.Coffe.Description() + ". Caramel"
}
