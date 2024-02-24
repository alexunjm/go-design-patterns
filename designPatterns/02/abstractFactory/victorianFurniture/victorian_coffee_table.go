package victorianfurniture

import "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"

type victorianCoffeeTable struct{}

func NewCoffeeTable() furniture.CoffeeTable {
	return &victorianCoffeeTable{}
}
