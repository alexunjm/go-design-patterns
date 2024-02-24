package factorymethod

import "github.com/alexunjm/go-design-patterns/designPatterns/01/factoryMethod/transport"

type seaLogistics struct {
}

// CreateTransport implements Logistics.
func (*seaLogistics) CreateTransport() transport.Transport {
	return transport.NewShip()
}

func NewSeaLogistics() Logistics {
	return &seaLogistics{}
}
