package AbstractFactoryPattern

import "fmt"

type Shirts interface {
	HasSleeves()
	SuitableFor()
}

type Pants interface {
	FullLength()
	SuitableFor()
}

type ClothesFactory interface {
	MakeShirts() Shirts
	MakePants() Pants
}

type WorkShirts struct{}

func (w WorkShirts) HasSleeves() {
	fmt.Println("Yes!! long sleeves.")
}

func (w WorkShirts) SuitableFor() {
	fmt.Println("Office wear.")
}

type WorkPants struct{}

func (w WorkPants) FullLength() {
	fmt.Println("Pants of 100 cm.")
}

func (w WorkPants) SuitableFor() {
	fmt.Println("Office wear.")
}

type VacationShirts struct{}

func (v VacationShirts) HasSleeves() {
	fmt.Println("No!! no sleeves.")
}

func (v VacationShirts) SuitableFor() {
	fmt.Println("Home wear.")
}

type VacationPants struct{}

func (v VacationPants) FullLength() {
	fmt.Println("Pants of 50 cm.")
}

func (v VacationPants) SuitableFor() {
	fmt.Println("Home wear.")
}

type OfficeFactory struct{}

func (o *OfficeFactory) MakeShirts() Shirts {
	return WorkShirts{}
}

func (o *OfficeFactory) MakePants() Pants {
	return WorkPants{}
}

type HomeFactory struct{}

func (h *HomeFactory) MakeShirts() Shirts {
	return VacationShirts{}
}

func (h *HomeFactory) MakePants() Pants {
	return VacationPants{}
}
