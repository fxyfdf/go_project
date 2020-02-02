package main

import "fmt"

//1 面向对象
/* 面向对象的好处：  封装继承多态
面向对象  继承  方法
面向过程（步骤过程） 每一步都是自己执行
面向对象（对象）：所谓的面向对象其实就是找一个专门做这个事的人来做，不用关心具体怎样实现的
所以说，面向过程强调的是过程，步骤。而面向对象强调的是对象，也就是干事的人 。我吃饭只关心饭，
不关心做饭的过程
类 ： 特征 行为 等 共同点构成了一类   具有相同属性的一类  如： 猫
*/
//1 对象创建   匿名字段实现继承以及对象创建
// go 语言中通过结构体表示一个类
type Student struct {
	// 属性 结构体成员
	// 方法  -- 函数
	//id int
	//name string
	//age int
	Person
	score float64
}

// 3 继承：  描述的是一种关系，
/*
继承是一种类之间的关系，描述一个类从另一个类获取成员信息的类间关系
继承必定发生在两个类之间，参与继承关系的双方成为父类和子类
父类提供成员信息，子类获取成员信息
*/
type Teacher struct {
	Person // 匿名字段
	salary float64
}
type Person struct {
	id   int
	name string
	age  int
}

// 5 指针类型匿名字段
type Student5 struct {
	*Person // 匿名字段
	score   float64
}

// 6 多重继承
type Object6 struct {
	id int
}
type Person6 struct {
	//id int
	Object6
	name string
	age  int
}
type Student6 struct {
	Person6 // 多重继承
	score   float64
}

// 7 为结构体添加方法
/*
func (对象 结构体类型) 方法名 （参数列表） (返回值列表){ 代码体 }

调用，1 实例化一个对象， 2 调用 对象名.方法名
*/
type Student7 struct {
	id   int
	name string
	age  int
}

func (s7 Student7) PrintShow() { //s7 相当于接收者
	fmt.Println("id=", s7.id)
	fmt.Println("name=", s7.name)
}
func (s7 *Student7) Editinfo() { // 接收者为指针时才能修改原始数据
	s7.id = 21
}

// 8 使用方法注意事项 ： 只要接收者类型不一样，这个方法就是同名，也是方法
// 接收者为制作类型
type Student8 struct {
	id   int
	name string
	age  int
}
type Teacher8 struct {
	id   int
	name string
}

func (s8 *Student8) Show() {
	fmt.Println(s8)
}
func (t8 *Teacher8) Show() {
	fmt.Println(t8)
}

func main() {
	/*
		//var stu Student=Student{1,"zs","18",88} // 类实例化
		var stu Student=Student{Person{12,"zs",18},88} // 类实例化,继承父类时大括号内独立初始化父类
		//var stu1 Student=Student{2,"le","18",88}
		var stu1 Student=Student{score: 88} //部分初始化
		var stu2 Student=Student{Person:Person{id:22}}  //部分初始化, 注意写法， Person:Persion{id:22}  对父类父类中的ID 初始化
		fmt.Println(stu)
		fmt.Println(stu1)
		fmt.Println(stu2)

		// 4 成员操作：获取成员值 对象名.成员 对象名.父类.成员   修改成员值    对象名.（父类.）成员= **  只修改对应对象的值
		fmt.Println(stu.age)  //  成员.属性
		stu.age = 89
		fmt.Println(stu.Person.age)  //  成员.父类.属性  ==  成员.属性

		// 5 指针类型匿名字段
		var stu5 Student5 = Student5{Person:&Person{
			id:   5,
			name: "王五",
			age:  95,
		}}
		fmt.Println(stu5.name) //打印时候需要通过对象名方式

		// 6 多重继承
		var stu6 Student6
		stu6.age = 18
		stu6.Person6.Object6.id = 12  //或 stu6.id  ，一层一层选择这个属性是否有， 减少多层继承的 书写
		fmt.Println("age=",stu6.age," ","id=",stu6.Person6.Object6.id)

		// 7 为结构体添加方法
		var stu7 Student7
		stu7 = Student7{
			id:   17,
			name: "liqi",
			age:  97,
		}
		stu7.PrintShow() // 完成对访问的调用
		stu7.Editinfo()
		stu7.PrintShow()
	*/
	//8
	stu8 := Student8{18, "liba", 18}
	t8 := Teacher8{id: 8, name: "t8"}
	stu8.Show() // ---> (&stu8.Show())
	t8.Show()
	//fmt.Println(stu8)
	//fmt.Println(t8)
}
