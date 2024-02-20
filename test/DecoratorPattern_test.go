package test

import (
	"DesignPatterns/DecoratorPattern"
	"fmt"
	"testing"
)

func TestDecoratorPattern(t *testing.T) {
	// 創建一個裝飾過的函數
	decorated := DecoratorPattern.Decorator(DecoratorPattern.Message)

	// 使用裝飾過的函數
	fmt.Println(decorated("world")) // 輸出: HELLO, WORLD
}
