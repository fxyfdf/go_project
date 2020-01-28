package main

import "fmt"

// 循环结构

func main() {

	/*	// for循环,基本语法如下
		var i int
		for i = 1; i < 2; i++ {
			fmt.Printf("%d\n", i)
		}

		// 求1-100 数字之和
		var num int = 1
		var sum int = 0
		for num = 1; num < 101; num++ {
			sum = sum + num
		}
		fmt.Printf("1-100 之和为：%d\n", sum)

		// 求1-100 偶数之和
		var num2 int = 0
		var sum2 int = 0
		for num2 = 0; num2 < 101; num2 +=2 {
			sum2 = sum2 + num2
		}
		fmt.Printf("1-100 偶数之和为：%d\n", sum2)

		// break结束当前循环   continue
		for i :=0; i<10; i++{
			if i==4 {
				break //i = 4 时结束当前循环,4 不会输出，后面不会执行
			}
			fmt.Println(i)
		}

		// break 应用
		// 要求用户输入用户名和密码， 只要不是admin，888888 就一直提示用户名密码错误
		// 1： 定义变量存储用户名密码
		var userName string
		var userPwd string
		var count int //记录输入次数
		for {  // 不加条件的死循环
		fmt.Println("请输入用户名：")
		fmt.Scan(&userName)
		fmt.Println("请输入密码：")
		fmt.Scan(&userPwd)
		// 2: 对输入的用户名密码进行判断
		if userName == "admin" && userPwd == "888888" {
			fmt.Println("输入正确")
			break
		} else {
			count++
			if count >=3 {
				fmt.Println("最多输错3次")
				break
			}
			// 3: 如果不正确，给出重新输入的提示，如果正确，给出正确提示
			fmt.Println("输入错误")
		}}

		// continue: 结束本轮循环，然后进行循环条件判断，如果成立进行，然后进行下一轮循环否则退出循环
		//  输出1-5  3 除外
		for i := 1; i<= 5; i++ {
			if i == 3 {
				continue //跳出此次循环，不执行此次循环的以后语句
			}
			fmt.Println("i = ",i)
		}
	*/

	// 循环嵌套
	// 打印五行星星
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
	// 打印水仙花数字  每位数字立方和等于本身
	var h int //百位数
	var t int //十位数
	var b int // 个位数
	for i := 100; i < 1000; i++ {
		//ctrl alt l   快速排版
		h = i / 100
		t = i % 100 / 10
		b = i % 10

		if i == h*h*h+t*t*t+b*b*b {
			fmt.Println("i = ", i)
		}
	}

}
