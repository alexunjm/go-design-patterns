package modernfurniture

import (
	"fmt"

	"github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"
)

type modernCoffeeTable struct{}

func NewCoffeeTable() furniture.CoffeeTable {
	fmt.Println("creating a new modern coffee table")
	return &modernCoffeeTable{}
}
