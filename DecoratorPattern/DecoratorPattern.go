package DecoratorPattern

import (
	"strings"
)

// FunctionType 定義一個函數類型
type FunctionType func(string) string

// Message 函數返回一個基本的消息
func Message(name string) string {
	return "Hello, " + name
}

// Decorator 函數是一個裝飾器，它接受一個 FunctionType 並返回一個新的 FunctionType
func Decorator(fn FunctionType) FunctionType {
	return func(name string) string {
		// 在調用原始函數之前或之後添加額外的行為
		return strings.ToUpper(fn(name))
	}
}
