package victorianfurniture

import "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"

type victorianSofa struct{}

func NewSofa() furniture.Sofa {
	return &victorianSofa{}
}
