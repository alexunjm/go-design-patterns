package factorymethod

import "github.com/alexunjm/go-design-patterns/designPatterns/01/factoryMethod/transport"

type SeaLogistics struct {
}

// CreateTransport implements Logistics.
func (*SeaLogistics) CreateTransport() transport.Transport {
	return transport.NewShip()
}

func NewSeaLogistics() Logistics {
	return &SeaLogistics{}
}
