package modernfurniture

import "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"

type modernChair struct{}

func NewChair() furniture.Chair {
	return &modernChair{}
}
