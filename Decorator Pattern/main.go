package main

import (
	"fmt"

	"github.com/mastering-go/decoration-pattern/coffe"
)

func main() {
	//
	sc := &coffe.SimpleCoffe{}

	coffeWithMilk := &coffe.Milk{Coffe: sc}

	fmt.Println("Coffe with milk: ", coffeWithMilk.Description(), "- Cost: ", coffeWithMilk.Cost())

	coffeWithCaramel := &coffe.Caramel{Coffe: sc}

	fmt.Println("Coffe with caramel: ", coffeWithCaramel.Description(), "- Cost: ", coffeWithCaramel.Cost())

	coffeWithMilkAndCaramel := &coffe.Milk{Coffe: &coffe.Caramel{Coffe: sc}}

	fmt.Println("Coffe with caramel and milk: ", coffeWithMilkAndCaramel.Description(), "- Cost: ", coffeWithMilkAndCaramel.Cost())

}
