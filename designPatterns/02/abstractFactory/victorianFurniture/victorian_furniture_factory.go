package victorianfurniture

import (
	abstractfactory "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory"
	"github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"
)

type VictorianFurnitureFactory struct{}

// CreateChair implements abstractfactory.FurnitureFactory.
func (*VictorianFurnitureFactory) CreateChair() furniture.Chair {
	return NewChair()
}

// CreateCoffeTable implements abstractfactory.FurnitureFactory.
func (*VictorianFurnitureFactory) CreateCoffeTable() furniture.CoffeeTable {
	return NewCoffeeTable()
}

// CreateSofa implements abstractfactory.FurnitureFactory.
func (*VictorianFurnitureFactory) CreateSofa() furniture.Sofa {
	return NewSofa()
}

func NewVictorianFurnitureFactory() abstractfactory.FurnitureFactory {
	return &VictorianFurnitureFactory{}
}
