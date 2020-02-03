package main

import "fmt"

// 面向对象案例

// 完成两个数加减的计算器

//1： 抽取类
//2 添加加减算法
//加法类
type Add struct {
	Object
}

// 减法类
type Sub struct {
	Object
}
type Object struct {
	num1 int
	num2 int
}

func (p *Add) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 + p.num2
}
func (p *Sub) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = 2
	return p.num1 - p.num2
}
func main() {
	var add Add
	n := add.GetResult(10, 5)
	fmt.Println(n)
	var sub Sub
	n2 := sub.GetResult(10, 5)
	fmt.Println(n2)
}
