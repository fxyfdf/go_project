package main

import "fmt"

// 1： 添加一个类
type Object struct {
}

// 2 添加一个方法
func (o *Object) GetResult(numA int, numB int, op string) int {
	//1 添加参数
	var result int
	//2 判断计算方式
	switch op {
	case "+":
		result = numA + numB
	case "-":
		result = numA - numB
	case "*":
		result = numA * numB
	case "/":
		result = numA / numB
	}
	return result
}

func main() {
	var obj Object
	num := obj.GetResult(10, 10, "+")
	fmt.Println("10+10=", num)
}
