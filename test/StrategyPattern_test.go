package test

import (
	"DesignPatterns/StrategyPattern"
	"testing"
)

func TestStrategyPattern(t *testing.T) {
	strategy := StrategyPattern.NewTransportation("scooter")

	if strategy == nil {
		t.Error("Invalid transportation strategy")
	}

	StrategyPattern.NavigateToTarget(strategy)
}
