package AdapterPattern

import "fmt"

type Charger interface {
	TypecCharge()
}

type User struct {
}

func (u *User) UseTypecCharge(charger Charger) {
	fmt.Println("The user uses TypeC cable to charge.")
	charger.TypecCharge()
}

type Android struct {
}

func (a *Android) TypecCharge() {
	fmt.Println("Charging with TypeC.")
}

type Iphone struct {
}

func (i *Iphone) LightningCharge() {
	fmt.Println("Charging with Lightning.")
}

type AppleAdapter struct {
	Device *Iphone
}

func (adapter *AppleAdapter) TypecCharge() {
	fmt.Println("Adapter converts Lightning to TypeC.")
	adapter.Device.LightningCharge()
}
