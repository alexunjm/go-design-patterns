package factorymethod_test

import (
	"testing"

	factorymethod "github.com/alexunjm/go-design-patterns/designPatterns/01/factoryMethod"
)

func Test_Logistics(t *testing.T) {
	testCases := []struct {
		desc      string
		logistics factorymethod.Logistics
	}{
		{
			desc:      "create transport for road",
			logistics: factorymethod.NewRoadLogistics(),
		},
		{
			desc:      "create transport for sea",
			logistics: factorymethod.NewSeaLogistics(),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			transport := tC.logistics.CreateTransport()
			transport.Deliver()
		})
	}
}
