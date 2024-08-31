package test

import (
	"DesignPatterns/AdapterPattern"
	"testing"
)

func TestAdapterPattern(t *testing.T) {
	user := &AdapterPattern.User{}

	dadsPhone := &AdapterPattern.Android{}
	user.UseTypecCharge(dadsPhone)

	myPhone := &AdapterPattern.Iphone{}
	appleDeviceAdapter := &AdapterPattern.AppleAdapter{
		Device: myPhone,
	}
	user.UseTypecCharge(appleDeviceAdapter)
}
