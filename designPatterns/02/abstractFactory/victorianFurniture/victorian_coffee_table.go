package victorianfurniture

import (
	"fmt"

	"github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"
)

type victorianCoffeeTable struct{}

func NewCoffeeTable() furniture.CoffeeTable {
	fmt.Println("creating a new victorian coffee table")
	return &victorianCoffeeTable{}
}
