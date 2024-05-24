package modernfurniture

import (
	"fmt"

	"github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"
)

type modernChair struct{}

func NewChair() furniture.Chair {
	fmt.Println("creating a new modern chair")
	return &modernChair{}
}
