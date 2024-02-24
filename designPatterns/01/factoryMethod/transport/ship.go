package transport

import "fmt"

type Ship struct{}

// Deliver implements Transport.
func (*Ship) Deliver() {
	fmt.Println("Delivering with Ship")
}

func NewShip() Transport {
	return &Ship{}
}
