package main

import "fmt"

func main() {

/*	// 变量初始化
	var age int
	var num int
	// 变量赋值
	var age1 int = 10
	var age2 int
	var age3 int
	age2 = 20
	age3 = age2
	// age3 = age2  // 原有的旧值被覆盖

	// 交换两个变量, 通过中间变量
	var num1 int = 10
	var num2 int = 20
	var tmp int
	tmp = num1
	num1 = num2
	num2 = tmp
	// 交换两个变量
	num10 := 10
	num20 := 20
	num10, num20 = num20, num10
	//fmt.Println(num10, num20)

	// 自动推倒类型
	num3 := 10
	// num4, num5, num6 := 1, 2, 3
	fmt.Println(num3)




	//声明变量必须使用，不使用时会报错Variable operating.go:13:6: age33 declared and not used
	fmt.Println(age1, age, age2, age3)
	fmt.Print(num ,"\n")*/

/*	//  输出格式
	var num int = 10
	fmt.Println("aaaaaa=", num)  // 输出后有回车
	fmt.Print("bbbbb=", num)
	fmt.Printf("%d", num)
	fmt.Printf("num = %d", num)
*/
	// 接收输入
	var age int
	fmt.Println("请输入年龄：")
	//fmt.Scanf("%d", &age) // 通过Scanf 函数将键盘的数据赋值给变量，但是变量前面加&
	//fmt.Println("age = ", age)
	//fmt.Println(&age)  //0x 开头16进制数据
	//fmt.Printf('%P', &age)
	fmt.Scan(&age)  //不需要指定地址符
	fmt.Println("age = ", age)
	fmt.Println(&age)  //0x 开头16进制数据

}
