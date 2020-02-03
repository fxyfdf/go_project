package main

import (
	"fmt"
)

//10 方法继承
type Student10 struct {
	Person10
	score float64
}
type Person10 struct {
	id   int
	name string
	age  int
}

func (p10 *Person10) PrintInfo10() {
	fmt.Println(*p10)

}

// 11 方法继承练习
/*
实现继承关系：
记者： 我叫张三，我的还好是摄鸟，我的年龄34，是一个男狗仔
程序员： 我叫孙全，我的年龄是23，性别男生，我的工作年限是3年、
怎样抽取类，成员，考虑是否有父类，从而实现成员，方法继承
父类 人 ： name age sex
子类 记者  程序员
*/

// 11.1 定义父类
type Person11 struct {
	name string
	age  int
	sex  string
}

// 11.2 给父类添加方法 初始化数据
func (p11 *Person11) SetValue11(userName string, userAge int, usreSex string) {
	p11.sex = usreSex
	p11.name = userName
	p11.age = userAge
}

// 11.3 定义子类
//记者
type Rep11 struct {
	Person11
	Hobby string
}

func (r11 *Rep11) RepSayHello(h string) {
	r11.Hobby = h
	fmt.Printf("我叫%s，我的还好是%s，我的年龄%d，是一个男狗仔\n", r11.name, r11.Hobby, r11.age)
}

//程序员
type Pro11 struct {
	Person11
	WorkYear int
}

func (p11 *Pro11) ProSayHello(wy int) {
	p11.WorkYear = wy
	fmt.Printf("我叫%s，我的年龄是%d，性别%s，我的工作年限是%d年\n", p11.name, p11.age, p11.sex, p11.WorkYear)
}

// 11.4 给子类添加方法

// 12 方法重写:  就是子类（结构体）中的方法，将父类中的相同名称的方法的功能重新给改写了（在调用时，默认调用的时子类中的方法）

//父类中的方法与子类方法重名，

type Person12 struct {
	name string
	age  int
}

func (p12 *Person12) PrintInfo() {
	fmt.Println("这是父类中的方法")
}

type Student12 struct {
	Person12
	score float64
}

func (p12 *Student12) PrintInfo() {
	fmt.Println("这是子类中的方法")
}

// 13 方法值与方法
type Person13 struct {
	name string
	age  int
}

func (p *Person13) PrintInfo13() {
	fmt.Println(*p)
}

// 14 方法总结
func main() {
	/*
		//子类
		stu := Student10{
			Person10: Person10{1,"aa",19},
			score:    99,
		}
		//调用父类继承
		stu.PrintInfo10()

		// 11 方法继承练习
		var rep Rep11
		rep.SetValue11("jizhe",22,"男")
		rep.RepSayHello("摄鸟")
		var pro Pro11
		// 调用公共属性，对成员进行初始化
		pro.SetValue11("cxy",23,"男")
		pro.ProSayHello(2)
	*/
	// 12
	var stu12 Student12
	stu12.PrintInfo()          //
	stu12.Person12.PrintInfo() //调用父类时需要写上父类
	//13
	per := Person13{"十三", 18}
	per.PrintInfo13()
	// 方法值
	f := per.PrintInfo13 //f 已经存在对象值
	fmt.Printf("%T\n", f)
	f()
	// 方法表达式
	f11 := (*Person13).PrintInfo13
	f11(&per)
}
