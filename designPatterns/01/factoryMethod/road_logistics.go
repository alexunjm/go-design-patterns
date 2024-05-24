package factorymethod

import "github.com/alexunjm/go-design-patterns/designPatterns/01/factoryMethod/transport"

type roadLogistics struct {
}

// CreateTransport implements Logistics.
func (*roadLogistics) CreateTransport() transport.Transport {
	return transport.NewTruck()
}

func NewRoadLogistics() Logistics {
	return &roadLogistics{}
}
