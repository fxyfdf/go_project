package main

import "fmt"

//17 接口：

/*
接口就是一种规范与标准，只是规定了要做那些事情。具体怎么做，接口是不管的。
接口把所有的具体共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
*/

// 18 接口定义方式

/*
接口定义语法：
type 接口名字 interface{
	方法声明// 只是声明，并未实现
}
接口实现：
*/

// 19 多态定义于是实现

/*
什么是多态：  指的是多种表现形式
多态就是同一个接口，使用不通的实例而执行不通操作

多态实现
func 函数名 (参数 接口类型){
}
*/

type Personer18 interface { // 定义接口，实现一个打招呼的功能
	SayHello() // 完成方法声明
}

type Student18 struct {
}

func (s *Student18) SayHello() {
	fmt.Println("老师好，我是学生")
}

type Teacher18 struct {
}

func (t18 *Teacher18) SayHello() {
	fmt.Println("学生好，我是老师")
}

// 多态，传输的参数不同，执行不通的方法
func WhoSayHi(h Personer18) {
	h.SayHello() // 接口的方法
}

// 20 多态案例

// 用多态来模拟实现 将移动硬盘或者U盘插到电脑上进行读写数据
// 定义一个接口
type Storager interface {
	Read()
	Write()
}

// 移动硬盘
type MDisk struct {
}

func (m MDisk) Read() {
	fmt.Println("移动硬盘读取数据")
}
func (m MDisk) Write() {
	fmt.Println("移动硬盘写入取数据")
}

//U盘
type UDisk struct {
}

func (u UDisk) Read() {
	fmt.Println("U盘读取数据")
}
func (u UDisk) Write() {
	fmt.Println("U盘写入数据")
}

// 定义一个函数，实现多态
func Computer(c Storager) {
	c.Read()
	c.Write()
}

// 21 objectDemo2

// 22 ObjectDemo3  计算器第二个版本

// 23 ObejctDemo4   计算器第三个版本

//24 多态总结

//25 接口继承与转换
type Humaner25 interface {
	SayHello25()
}
type Personer25 interface {
	Humaner25
	Say25()
}
type Student25 struct {
}

func (s *Student25) SayHello25() {
	fmt.Println("大家好")
}
func (s *Student25) Say25() {
	fmt.Println("你好")
}

//26 空接口定义与使用 main 中

//27 类型断言

/*
通过类型断言，可以判断空接口中存储的数据类型。
语法:
value, ok :=m.(T)
m: 表空接口类型变量
T：是断言的类型
value: 变量m中的值
ok: 布尔类型变量，如果断言成功为true否则 false
*/

// 28 空接口与类型断言综合应用  ObjectDemo5

func main() {
	/*
		// 18

		// 对象名.方法名
		var stu18 Student18
		stu18.SayHello()
		var t18 Teacher18
		t18.SayHello()
		// 通过接口变量来调用,必须都实现接口中的方法
		var person18 Personer
		person18=&stu18
		person18.SayHello()

		// 19
		var stu18 Student18
		var t19 Teacher18
		WhoSayHi(&stu18) // 注意取地址符号
		WhoSayHi(&t19)

		// 20 多态案例
		var mdisk MDisk
		var udisk UDisk
		// 多态方式调用
		Computer(&mdisk)
		Computer(&udisk)
		// 接口方式调用
		var stor Storager // 实例化接口
		stor = &mdisk // 执行具体实例
		stor.Write() // 执行实例方法
		stor.Read()

		//25
		var stu25 Student25
		var per25 Personer25
		per25 = &stu25
		per25.Say25()
		per25.SayHello25()

		var h Humaner25
		h = per25
		h.SayHello25()

	*/
	// 26 空接口定义与使用
	// 没有任何方法的接口，可以存储任何数据
	var i interface{}
	i = 123
	i = "abc"
	fmt.Println(i)
	var s []interface{}
	s = append(s, "123", "abd")
	for j := 0; j < len(s); j++ {
		fmt.Println(s[j])
	}
	// 27 类型断言
	var i27 interface{}
	i27 = 213
	i27 = "ddddd"
	value, ok := i27.(int)
	if ok {
		fmt.Println(value + 10)
	} else {
		fmt.Println("类型推断错误")
	}
}
