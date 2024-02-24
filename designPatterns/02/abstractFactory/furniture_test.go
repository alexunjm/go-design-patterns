package abstractfactory_test

import (
	"fmt"
	"testing"

	abstractfactory "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory"
	modernfurniture "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/modernFurniture"
	victorianfurniture "github.com/alexunjm/go-design-patterns/designPatterns/02/abstractFactory/victorianFurniture"
)

type testCase struct {
	desc    string
	factory abstractfactory.FurnitureFactory
}

func Test_FurnitureFactory(t *testing.T) {
	testCases := []testCase{
		{
			desc:    "victorian",
			factory: victorianfurniture.NewVictorianFurnitureFactory(),
		},
		{
			desc:    "modern",
			factory: modernfurniture.NewModernFurnitureFactory(),
		},
	}
	for _, tC := range testCases {
		FurnitureTestsUsing(t, &tC)
	}
}

type furnitureTest struct {
	factory abstractfactory.FurnitureFactory
}

func (t *furnitureTest) testChair(_ *testing.T) {
	t.factory.CreateChair()
}
func (t *furnitureTest) testCoffeeTable(_ *testing.T) {
	t.factory.CreateCoffeTable()
}
func (t *furnitureTest) testSofa(_ *testing.T) {
	t.factory.CreateSofa()
}

func FurnitureTestsUsing(t *testing.T, tc *testCase) {
	test := &furnitureTest{tc.factory}

	testCases := []struct {
		desc   string
		method func(_ *testing.T)
	}{
		{
			desc:   fmt.Sprintf("%s %s", tc.desc, "chair"),
			method: test.testChair,
		},
		{
			desc:   fmt.Sprintf("%s %s", tc.desc, "coffe table"),
			method: test.testCoffeeTable,
		},
		{
			desc:   fmt.Sprintf("%s %s", tc.desc, "sofa"),
			method: test.testSofa,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, tC.method)
	}
}
