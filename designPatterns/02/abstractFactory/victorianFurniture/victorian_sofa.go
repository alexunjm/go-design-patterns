package victorianfurniture

import (
	"fmt"

	"github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/furniture"
)

type victorianSofa struct{}

func NewSofa() furniture.Sofa {
	fmt.Println("creating a new victorian sofa")
	return &victorianSofa{}
}
