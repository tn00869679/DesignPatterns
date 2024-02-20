package test

import (
	"DesignPatterns/AbstractFactoryPattern"
	"fmt"
	"testing"
)

func TestAbstractFactoryPattern(t *testing.T) {
	factory := getCoffeeFactory()

	sugar := factory.AddSugar()
	milk := factory.AddMilk()

	sugar.SetSugarCategory()
	milk.SetMilkCategory()

	fmt.Println("A cup of ", factory.GetDescription())
}

func getCoffeeFactory() AbstractFactoryPattern.CoffeeFactory {
	return AbstractFactoryPattern.StarbucksFactory{}
}
