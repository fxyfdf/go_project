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
func (a *Add) SetData(data ...interface{}) bool {
	//1 对数据的个数进行校验
	var b bool = true
	if len(data) > 2 {
		b = false
		fmt.Println("参数格式错误")
	}
	value, ok := data[0].(int) //类型断言，判断空接口中数据类型是否正确
	if !ok {
		fmt.Println("第一个数类型错误")
		b = false
	}
	value1, ok1 := data[1].(int) //类型断言，判断空接口中数据类型是否正确
	if !ok1 {
		fmt.Println("第一个数类型错误")
		b = false
	}
	//2 类型进行校验: 通过校验后，把值传入
	a.numA = value
	a.NumB = value1
	return b
}

// 减法
type Sub struct {
	Object
}

func (a *Sub) GetResult() int { //方法的实现要和接口中方法的声明保持一致
	return a.numA - a.NumB
}
func (a *Sub) SetData(data ...interface{}) bool {
	//1 对数据的个数进行校验
	var b bool = true
	if len(data) > 2 {
		b = false
		fmt.Println("参数格式错误")
	}
	value, ok := data[0].(int) //类型断言，判断空接口中数据类型是否正确
	if !ok {
		fmt.Println("第一个数类型错误")
		b = false
	}
	value1, ok1 := data[1].(int) //类型断言，判断空接口中数据类型是否正确
	if !ok1 {
		fmt.Println("第一个数类型错误")
		b = false
	}
	//2 类型进行校验: 通过校验后，把值传入
	a.numA = value
	a.NumB = value1
	return b
}

type Object struct {
	numA int
	NumB int
}
type Resulter interface {
	GetResult() int // 返回类型
	// 三个点，不定参数列表，可以传入多个参数
	SetData(data ...interface{}) bool // 完成参与计算的数据类型校验
}

// 对象问题
// 1 定义一个新的类
type OperatorFactory struct {
}

// 2 创建一个方法，在该方法中完成对象的创建  （封装）
func (o *OperatorFactory) CreateOperator(op string) Resulter {
	switch op {
	case "+":
		add := new(Add) // 指定类型开辟一个地址空间
		return add
	case "-":
		sub := new(Sub)
		return sub
	default:
		return nil
	}

}

// 多态
func OperatotWho(h Resulter) int {
	return h.GetResult()
}

func main() {
	var operator OperatorFactory
	obj := operator.CreateOperator("-")
	b := obj.SetData(30, 10)
	if b {
		num := OperatotWho(obj)
		fmt.Println(num)

	}
}
