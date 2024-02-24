package victorianfurniture

import "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"

type victorianChair struct{}

func NewChair() furniture.Chair {
	return &victorianChair{}
}
