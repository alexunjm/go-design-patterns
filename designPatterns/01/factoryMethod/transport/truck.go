package transport

import "fmt"

type Truck struct{}

// Deliver implements Transport.
func (*Truck) Deliver() {
	fmt.Println("Delivering with Truck")
}

func NewTruck() Transport {
	return &Truck{}
}
