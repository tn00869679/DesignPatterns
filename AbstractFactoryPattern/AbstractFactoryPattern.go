package AbstractFactoryPattern

import "fmt"

type Sugar interface {
	SetSugarCategory()
}

type Milk interface {
	SetMilkCategory()
}

type CoffeeFactory interface {
	AddSugar() Sugar
	AddMilk() Milk
	GetDescription() string
}

type FamilySugar struct{}

func (oneself FamilySugar) SetSugarCategory() {
	fmt.Println("Japan's sugar")
}

type FamilyMilk struct{}

func (oneself FamilyMilk) SetMilkCategory() {
	fmt.Println("Japan's milk")
}

type SevenSugar struct{}

func (oneself SevenSugar) SetSugarCategory() {
	fmt.Println("Taiwan's sugar")
}

type SevenMilk struct{}

func (oneself SevenMilk) SetMilkCategory() {
	fmt.Println("Taiwan's milk")
}

type StarbucksSugar struct{}

func (oneself StarbucksSugar) SetSugarCategory() {
	fmt.Println("American's sugar")
}

type StarbucksMilk struct{}

func (oneself StarbucksMilk) SetMilkCategory() {
	fmt.Println("American's milk")
}

type FamilyFactory struct{}

func (oneself FamilyFactory) AddSugar() Sugar {
	return FamilySugar{}
}

func (oneself FamilyFactory) AddMilk() Milk {
	return FamilyMilk{}
}

func (oneself FamilyFactory) GetDescription() string {
	return "Let's cafe"
}

type SevenFactory struct{}

func (oneself SevenFactory) AddSugar() Sugar {
	return SevenSugar{}
}

func (oneself SevenFactory) AddMilk() Milk {
	return SevenMilk{}
}

func (oneself SevenFactory) GetDescription() string {
	return "City cafe"
}

type StarbucksFactory struct{}

func (oneself StarbucksFactory) AddSugar() Sugar {
	return StarbucksSugar{}
}

func (oneself StarbucksFactory) AddMilk() Milk {
	return StarbucksMilk{}
}

func (oneself StarbucksFactory) GetDescription() string {
	return "Starbucks"
}
