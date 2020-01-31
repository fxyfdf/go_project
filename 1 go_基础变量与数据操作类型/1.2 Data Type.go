package main

import "fmt"

func main() {
	// 基本数据类型
	// 1 整型： 有符号 无符号
	var age uint
	age = 10
	fmt.Print(age, "\n")

	// 2浮点类型： float32  float64  通常使用float64
	/*	var num float64
		num = 1.23
		num2 := 11.11  //自动推导类型，生成的类型为float64
		fmt.Printf("%f\n", num)
		fmt.Printf("%.2f", num)  // 保留小数点后两位
		fmt.Printf("%T\n", num2)  // T输出变量类型
		fmt.Printf("%f\n", num2)  // 保留小数点后两位*/

	// 3 bool 类型
	/*	var isResult bool
		fmt.Println(isResult)
		fmt.Printf("%t", isResult)  // t 输出bool 类型，默认值为false*/

	// 4 字符类型
	/*var ch byte = 'a' // 单引号引起来的单个字符
	fmt.Println(ch)  //ASCII  大小写相差32
	fmt.Printf("ch = '%c'\n", ch)  // 输出字符， %c
	fmt.Printf("ASCII(ch) = %d", ch)*/

	// 5 字符串类型 : 定义(双引号引起来的多个字符)
	/*	var name string = "fxyfdf"
		fmt.Printf("name = %s \n", name)
		str := "a"
		fmt.Printf("str = %s\n", str)  // 字符隐藏结束符号  \0
		fmt.Printf("type(str) = %T \n",str)
		fmt.Printf("len(name) = %d \n", len(name) )
		fmt.Printf("len(汉字) = %d \n", len("汉字") ) // 一个汉字占用4个字符*/

	// 6 强制类型转换
	/*	var num float64 = 3.15
		fmt.Printf("num = %d\n", num)
		fmt.Printf("int(num) = %d\n", int(num))

		var num2 float64 = 5.9
		var num3 float64 = 6.62222
		fmt.Printf("num2 = %.2f\n", num2) //
		fmt.Printf("num3 = %.2f", num3)*/
	// fmt 包 格式化输入与输出
	// %c  ---字符类型 %d-数字  %f-浮点数字 %p-内存地址 %s-字符串 %t-输出布尔类型 %T-输出值类型
	/*	num := 11
		fmt.Printf("%b \n", num) // %b 表示二进制
		fmt.Printf("%o\n", num)  // %o 表示8进制
		fmt.Printf("%x\n", num)  // %x 表示16进制
		fmt.Printf("%X\n", num)  // %x 表示16进制*/

	//综合使用案例：
	// 请输入姓名，年龄，当前用户输入完以后在屏幕上显示: 您好xx,您的年龄是
	// 定义变量
	var name string
	var age1 int
	//1 请输入您的姓名
	fmt.Printf("请输入您的姓名：\n")
	fmt.Scan(&name)
	// 2 请输入您的年龄
	fmt.Printf("请输入您的年龄：\n")
	fmt.Scan(&age1)
	// 3 输出
	fmt.Printf("您好 %s,您的你年龄是 %d", name, age1)
}
