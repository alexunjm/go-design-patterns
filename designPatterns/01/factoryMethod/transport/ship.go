package transport

import "fmt"

type ship struct{}

// Deliver implements Transport.
func (*ship) Deliver() {
	fmt.Println("Delivering with Ship")
}

func NewShip() Transport {
	return &ship{}
}
