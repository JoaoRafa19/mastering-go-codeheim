package coffe

// Concrete Component
type SimpleCoffe struct {
}

func (s *SimpleCoffe) Cost() float64 {
	return 2.0

}
func (s *SimpleCoffe) Description() string {
	return "Simple Coffe"
}
