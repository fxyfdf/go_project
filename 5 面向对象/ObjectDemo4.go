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

// 集成
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

// 对象问题
// 1 定义一个新的类
type OperatorFactory struct {
}

// 2 创建一个方法，在该方法中完成对象的创建  （封装）
func (o *OperatorFactory) CreateOperator(op string, numA int, NumB int) int {
	switch op {
	case "+":
		add := Add{Object{numA, NumB}}
		return OperatotWho(&add)
	case "-":
		sub := Sub{Object{numA, NumB}}
		return OperatotWho(&sub)
	default:
		return 0
	}

}

// 多态
func OperatotWho(h Resulter) int {
	return h.GetResult()
}

func main() {
	var operator OperatorFactory
	aa := operator.CreateOperator("+", 10, 10)
	fmt.Println(aa)
}
