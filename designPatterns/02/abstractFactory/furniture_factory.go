package abstractfactory

import "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"

type FurnitureFactory interface {
	CreateChair() furniture.Chair
	CreateCoffeTable() furniture.CoffeeTable
	CreateSofa() furniture.Sofa
}
