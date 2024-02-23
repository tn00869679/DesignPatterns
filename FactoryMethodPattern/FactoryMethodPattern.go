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

func (t *Tea) GetSweetnessLevel() string {
	return fmt.Sprintf("%d percent sugar", t.sweetnessLevel)
}

func (t *Tea) SetSweetnessLevel(percent int) {
	t.sweetnessLevel = percent
}

func (t *Tea) HasCaffeine() bool {
	return t.hasCaffeine
}

func (t *Tea) SetCaffeine(hasCaffeine bool) {
	t.hasCaffeine = hasCaffeine
}

func (t *Tea) GetDescription() string {
	return "Green tea"
}

type Coffee struct {
	sweetnessLevel int
	hasCaffeine    bool
}

func (c *Coffee) GetSweetnessLevel() string {
	return fmt.Sprintf("%d percent sugar", c.sweetnessLevel)
}

func (c *Coffee) SetSweetnessLevel(percent int) {
	c.sweetnessLevel = percent
}

func (c *Coffee) HasCaffeine() bool {
	return c.hasCaffeine
}

func (c *Coffee) SetCaffeine(hasCaffeine bool) {
	c.hasCaffeine = hasCaffeine
}

func (c *Coffee) GetDescription() string {
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
