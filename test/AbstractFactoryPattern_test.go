package test

import (
	"DesignPatterns/AbstractFactoryPattern"
	"fmt"
	"testing"
)

func TestAbstractFactoryPattern(t *testing.T) {
	var factory AbstractFactoryPattern.ClothesFactory

	clothesStyle := buyClothes()
	switch clothesStyle {
	case "Japan":
		factory = &AbstractFactoryPattern.JapanFactory{}
	case "Taiwan":
		factory = &AbstractFactoryPattern.TaiwanFactory{}
	default:
		panic("Undefined clothes style")
	}

	clothes := factory.MakeClothes()
	clothes.Prepare()
	clothes.Tailor()

	fmt.Println("Clothes type :", factory.GetDescription())
}

func buyClothes() string {
	return "Taiwan"
}
