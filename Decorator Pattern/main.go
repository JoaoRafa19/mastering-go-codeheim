package main

import "fmt"

// Component interface
type Coffe interface {
	Cost() float64
	Description() string
}

// Concrete Component
type SimpleCoffe struct {
}

func (s *SimpleCoffe) Cost() float64 {
	return 2.0

}
func (s *SimpleCoffe) Description() string {
	return "Simple Coffe"
}

type Milk struct {
	coffe Coffe
}

func (m *Milk) Cost() float64 {
	return m.coffe.Cost() + 0.5
}

func (m *Milk) Description() string {
	return m.coffe.Description() + ". Milk"
}

type Caramel struct {
	coffe Coffe
}

func (c *Caramel) Cost() float64 {
	return c.coffe.Cost() + 1
}

func (c *Caramel) Description() string {
	return c.coffe.Description() + ". Caramel"
}

func main() {
	//
	coffe := &SimpleCoffe{}

	coffeWithMilk := &Milk{coffe: coffe}

	fmt.Println("Coffe with milk: ", coffeWithMilk.Description(), "- Cost: ", coffeWithMilk.Cost())

	coffeWithCaramel := &Caramel{coffe: coffe}

	fmt.Println("Coffe with caramel: ", coffeWithCaramel.Description(), "- Cost: ", coffeWithCaramel.Cost())

	coffeWithMilkAndCaramel := &Milk{coffe: &Caramel{coffe: coffe}}

	fmt.Println("Coffe with caramel and milk: ", coffeWithMilkAndCaramel.Description(), "- Cost: ", coffeWithMilkAndCaramel.Cost())

}
