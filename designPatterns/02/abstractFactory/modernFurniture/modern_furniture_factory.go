package modernfurniture

import (
	abstractfactory "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory"
	"github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"
)

type modernFurnitureFactory struct{}

// CreateChair implements abstractfactory.FurnitureFactory.
func (*modernFurnitureFactory) CreateChair() furniture.Chair {
	return NewChair()
}

// CreateCoffeTable implements abstractfactory.FurnitureFactory.
func (*modernFurnitureFactory) CreateCoffeTable() furniture.CoffeeTable {
	return NewCoffeeTable()
}

// CreateSofa implements abstractfactory.FurnitureFactory.
func (*modernFurnitureFactory) CreateSofa() furniture.Sofa {
	return NewSofa()
}

func NewModernFurnitureFactory() abstractfactory.FurnitureFactory {
	return &modernFurnitureFactory{}
}
