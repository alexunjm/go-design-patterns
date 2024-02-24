package factorymethod

import "github.com/alexunjm/go-design-patterns/designPatterns/01/factoryMethod/transport"

type Logistics interface {
	CreateTransport() transport.Transport
}
