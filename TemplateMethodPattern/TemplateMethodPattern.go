package TemplateMethodPattern

import (
	"fmt"
)

type SpendMoney interface {
	goToStore()
	buySomething()
	pay()
	goHome()
}

type Shopping struct {
}

func (s *Shopping) goToStore() {
}

func (s *Shopping) buySomething() {
}

func (s *Shopping) pay() {
	fmt.Println("Pay money.")
}

func (s *Shopping) goHome() {
	fmt.Println("Go home.")
}

func GoShopping(shop SpendMoney) {
	shop.goToStore()
	shop.buySomething()
	shop.pay()
	shop.goHome()
}

type Shoes struct {
	Shopping
}

func (s *Shoes) goToStore() {
	fmt.Println("Go to Nike.")
}

func (s *Shoes) buySomething() {
	fmt.Println("Buy the slippers.")
}

type Handbag struct {
	Shopping
}

func (h *Handbag) goToStore() {
	fmt.Println("Go to the department store.")
}

func (h *Handbag) buySomething() {
	fmt.Println("Buy the GUCCI.")
}
