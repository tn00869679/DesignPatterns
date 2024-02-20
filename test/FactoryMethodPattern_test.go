package test

import (
	"DesignPatterns/FactoryMethodPattern"
	"fmt"
	"testing"
)

func TestFactoryMethodPattern(t *testing.T) {
	beverageASweetnessLevel := 30
	beverageA := FactoryMethodPattern.BeverageFactory("tea", beverageASweetnessLevel, false)
	if beverageA == nil {
		fmt.Println("Unknown beverageA type")
	} else {
		fmt.Printf("Ordered: %s\n", beverageA.GetDescription())
		fmt.Println("SweetnessLevel: ", beverageA.GetSweetnessLevel())
		fmt.Println("Caffeine: ", beverageA.HasCaffeine())
	}

	beverageBSweetnessLevel := 50
	beverageB := FactoryMethodPattern.BeverageFactory("coffee", beverageBSweetnessLevel, true)
	if beverageB == nil {
		fmt.Println("Unknown beverageB type")
	} else {
		fmt.Printf("Ordered: %s\n", beverageB.GetDescription())
		fmt.Println("SweetnessLevel: ", beverageB.GetSweetnessLevel())
		fmt.Println("Caffeine: ", beverageB.HasCaffeine())
	}
}
