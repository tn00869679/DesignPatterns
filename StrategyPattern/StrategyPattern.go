package StrategyPattern

import "fmt"

type Transportation interface {
	travel()
}

type Car struct {
}

func (c *Car) travel() {
	fmt.Println("Travel by car")
}

type Scooter struct {
}

func (s *Scooter) travel() {
	fmt.Println("Travel by scooter")
}

type Train struct {
}

func (t *Train) travel() {
	fmt.Println("Travel by train")
}

func NewTransportation(strategyType string) Transportation {
	switch strategyType {
	case "car":
		return &Car{}
	case "scooter":
		return &Scooter{}
	case "train":
		return &Train{}
	default:
		return nil
	}
}

func NavigateToTarget(strategy Transportation) {
	strategy.travel()
}
