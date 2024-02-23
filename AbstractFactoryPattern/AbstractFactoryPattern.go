package AbstractFactoryPattern

import "fmt"

type Clothes interface {
	Prepare()
	Tailor()
}

type ClothesFactory interface {
	MakeClothes() Clothes
	GetDescription() string
}

type KimonoClothes struct{}

func (k KimonoClothes) Prepare() {
	fmt.Println("Prepare the kimono materials")
}

func (k KimonoClothes) Tailor() {
	fmt.Println("Just tailor with Japan's technique")
}

type FiftyPercentClothes struct{}

func (f FiftyPercentClothes) Prepare() {
	fmt.Println("Prepare the common materials of Taiwan")
}

func (f FiftyPercentClothes) Tailor() {
	fmt.Println("Just tailor with Taiwan's technique")
}

type JapanFactory struct{}

func (j *JapanFactory) MakeClothes() Clothes {
	return KimonoClothes{}
}

func (j *JapanFactory) GetDescription() string {
	return "Kimono"
}

type TaiwanFactory struct{}

func (t *TaiwanFactory) MakeClothes() Clothes {
	return FiftyPercentClothes{}
}

func (t *TaiwanFactory) GetDescription() string {
	return "Fifty Percent"
}
