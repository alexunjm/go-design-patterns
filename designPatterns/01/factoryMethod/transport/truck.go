package transport

import "fmt"

type truck struct{}

// Deliver implements Transport.
func (*truck) Deliver() {
	fmt.Println("Delivering with Truck")
}

func NewTruck() Transport {
	return &truck{}
}
