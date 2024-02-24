package modernfurniture

import "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"

type modernCoffeeTable struct{}

func NewCoffeeTable() furniture.CoffeeTable {
	return &modernCoffeeTable{}
}
