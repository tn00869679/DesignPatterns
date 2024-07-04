package test

import (
	"DesignPatterns/TemplateMethodPattern"
	"testing"
)

func TestTemplateMethodPattern(t *testing.T) {
	shoes := &TemplateMethodPattern.Shoes{}
	TemplateMethodPattern.GoShopping(shoes)

	handbag := &TemplateMethodPattern.Handbag{}
	TemplateMethodPattern.GoShopping(handbag)
}
