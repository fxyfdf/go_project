package main

import "fmt"

func main() {
	/*//算术运算符 + - * / %(取余) ++（自增) --（）
	var num1 int = 20
	var num2 int = 10
	fmt.Println("num1 + num2 = ", num1 + num2)
	// ++ --只能写在变量名的后面
	num1++
	fmt.Println("num1 + 1 = ", num1)
	num1--
	fmt.Println("num1 - 1 = ", num1)
	*/

	/*
	// 计算 语文，数学，英语  总分，平均分
	// 1 定义变量存储分数
	var chiness int = 90
	var math int = 80
	var english int = 69
	// 2 计算总分
	var score int
	score = chiness + math + english
	// 3 计算平均分
	var age float64
	age = float64(score)/3
	fmt.Printf("总分：%d \n", score)
	fmt.Printf("平均分：%.3f \n", age)
	*/

	// 赋值运算符
	var num int = 10
	num += 10 // num = num +10
	num -= 5  //num = num - 5
	num %= 2 + 3 //num %=5 num = num%5

	// 关系运算符:  < > >= <= ==   结果为 bool类型， True 或false

	// 逻辑运算符:   !非  && 与（都真为真）  || 或 一真为真

	// 运算符优先级
	/*
	特殊运算符： （）括号最高， . 包名.函数
	单目运算符： ！ 逻辑非  & 取地址
	双目运算符
	1. 算术运算符 * / %
	2. 算术运算符 + -
	3. 比较运算符 < > <= >= == !=
	4.逻辑运算符 &&
	5. 逻辑运算符 ||
	6. 赋值运算符 = += -= *= /=

	 */
	// 判断一年是不是闰年

	// 1声明一份变量存放数字
	var year int
	fmt.Printf("输入年份:")
	fmt.Scan(&year)
	// 2 能否被400 整除
	// 3 年份能够被4整除但不能被100 整除
	a := year % 400 ==0 || year % 4 == 0 && year % 100 != 0
	fmt.Println(a)
}
