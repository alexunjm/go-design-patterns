package modernfurniture

import (
	"fmt"

	"github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"
)

type modernSofa struct{}

func NewSofa() furniture.Sofa {
	fmt.Println("creating a new modern sofa")
	return &modernSofa{}
}
