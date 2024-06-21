package test

import (
	"DesignPatterns/AbstractFactoryPattern"
	"testing"
)

func TestAbstractFactoryPattern(t *testing.T) {
	var factory AbstractFactoryPattern.ClothesFactory

	clothesStyle := buyClothes()
	switch clothesStyle {
	case "Office":
		factory = &AbstractFactoryPattern.OfficeFactory{}
	case "Home":
		factory = &AbstractFactoryPattern.HomeFactory{}
	default:
		panic("Undefined clothes style")
	}

	shirts := factory.MakeShirts()
	shirts.HasSleeves()
	shirts.SuitableFor()

	pants := factory.MakePants()
	pants.FullLength()
	pants.SuitableFor()
}

func buyClothes() string {
	return "Home"
}
