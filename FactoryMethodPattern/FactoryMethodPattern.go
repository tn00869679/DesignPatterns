package FactoryMethodPattern

import "fmt"

type Beverage interface {
	GetSweetnessLevel() string
	SetSweetnessLevel(percent int)
	HasCaffeine() bool
	SetCaffeine(hasCaffeine bool)
	GetDescription() string
}

type Tea struct {
	sweetnessLevel int
	hasCaffeine    bool
}

func (beverage *Tea) GetSweetnessLevel() string {
	return fmt.Sprintf("%d percent sugar", beverage.sweetnessLevel)
}

func (beverage *Tea) SetSweetnessLevel(percent int) {
	beverage.sweetnessLevel = percent
}

func (beverage *Tea) HasCaffeine() bool {
	return beverage.hasCaffeine
}

func (beverage *Tea) SetCaffeine(hasCaffeine bool) {
	beverage.hasCaffeine = hasCaffeine
}

func (beverage *Tea) GetDescription() string {
	return "Green tea"
}

type Coffee struct {
	sweetnessLevel int
	hasCaffeine    bool
}

func (beverage *Coffee) GetSweetnessLevel() string {
	return fmt.Sprintf("%d percent sugar", beverage.sweetnessLevel)
}

func (beverage *Coffee) SetSweetnessLevel(percent int) {
	beverage.sweetnessLevel = percent
}

func (beverage *Coffee) HasCaffeine() bool {
	return beverage.hasCaffeine
}

func (beverage *Coffee) SetCaffeine(hasCaffeine bool) {
	beverage.hasCaffeine = hasCaffeine
}

func (beverage *Coffee) GetDescription() string {
	return "Latte"
}

func BeverageFactory(beverageType string, sweetnessLevel int, hasCaffeine bool) Beverage {
	var beverage Beverage
	switch beverageType {
	case "tea":
		beverage = &Tea{}
	case "coffee":
		beverage = &Coffee{}
	default:
		return nil
	}

	beverage.SetSweetnessLevel(sweetnessLevel)
	beverage.SetCaffeine(hasCaffeine)
	return beverage
}
