package main

import (
	"errors"
	"fmt"
)

var DivisionByZero = errors.New("除0错误")
var LessThanZero = errors.New("小于0错误")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionByZero
	}
	if a < 0 {
		return 0, LessThanZero
	}
	return a / b, nil
}

func main() {
	result, err := divide(-1, 10)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
