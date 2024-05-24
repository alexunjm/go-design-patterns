package victorianfurniture

import (
	"fmt"

	"github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"
)

type victorianChair struct{}

func NewChair() furniture.Chair {
	fmt.Println("creating a new victorian chair")
	return &victorianChair{}
}
