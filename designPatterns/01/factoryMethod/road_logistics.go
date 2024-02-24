package factorymethod

import "github.com/alexunjm/go-design-patterns/designPatterns/01/factoryMethod/transport"

type RoadLogistics struct {
}

// CreateTransport implements Logistics.
func (*RoadLogistics) CreateTransport() transport.Transport {
	return transport.NewTruck()
}

func NewRoadLogistics() Logistics {
	return &RoadLogistics{}
}
