package main

import (
	factorymethod "github.com/alexunjm/go-design-patterns/designPatterns/01/factoryMethod"
	abstractfactory "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory"
	modernfurniture "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/modernFurniture"
	victorianfurniture "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/victorianFurniture"
)

func main() {
	// runFactoryMethod()
	runAbstractFactory()
}

func runAbstractFactory() {
	factories := []abstractfactory.FurnitureFactory{
		victorianfurniture.NewVictorianFurnitureFactory(),
		modernfurniture.NewModernFurnitureFactory(),
	}
	for _, factory := range factories {
		factory.CreateChair()
		factory.CreateCoffeTable()
		factory.CreateSofa()
	}
}

func runFactoryMethod() {
	logistics := []factorymethod.Logistics{
		factorymethod.NewRoadLogistics(),
		factorymethod.NewSeaLogistics(),
	}
	for _, l := range logistics {
		l.CreateTransport().Deliver()
	}
}
