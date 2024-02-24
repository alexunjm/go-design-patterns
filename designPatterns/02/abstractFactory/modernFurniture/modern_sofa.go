package modernfurniture

import "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"

type modernSofa struct{}

func NewSofa() furniture.Sofa {
	return &modernSofa{}
}
