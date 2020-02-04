package main

import "fmt"

// 实现要给计算器，增加多态

// 加法类
type Add struct {
	Object
}

func (a *Add) GetResult() int { //方法的实现要和接口中方法的声明保持一致
	return a.numA + a.NumB
}

type Sub struct {
	Object
}

func (a *Sub) GetResult() int { //方法的实现要和接口中方法的声明保持一致
	return a.numA - a.NumB
}

type Object struct {
	numA int
	NumB int
}
type Resulter interface {
	GetResult() int // 返回类型
}

func main() {
	add := Add{Object{10, 20}}
	var add2 Add
	add2 = Add{Object{20, 20}} // Object需要写

	numa := add.GetResult()
	numa2 := add2.GetResult()
	fmt.Println(numa)
	fmt.Println(numa2)
}
