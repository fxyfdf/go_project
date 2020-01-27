package main

import "fmt"

func main() {
	//分支结构： 执行到某一地方会选择执行一个选择判断
	// if swith
	/*
	fmt.Println("请输入您的年龄：")
	var score int
	fmt.Scan(&score)
	if score >= 18 &&  score < 35{
		fmt.Println("您成年了18-35")
	}
	fmt.Println("你好 if ") //条件外
	if score >=18 && score < 35{
		fmt.Println("您成年了18-35")
	} else {
		fmt.Println("未成年或大于35")
	}
	fmt.Println("你好 if  else") //条件外
*/

	//判断一个数字是不是偶数
	var num int
	fmt.Printf("输入一个数字：")
	fmt.Scan(&num)
	if num % 2 == 0 {
		fmt.Println(num,"是偶数 \n")
		if num > 100 {
			fmt.Printf("%d 大于100 了\n", num) //多层嵌套
		}
	} else {
		fmt.Println(num,"是奇数 \n")
	}

	//对学员成绩进行评测
	fmt.Printf("输入学员成绩：")
	var score int
	var grade string
	fmt.Scan(&score)
	if score >= 90{
		grade = "A"
		fmt.Printf("A")
	} else if score >= 80 {
		grade = "B"
		fmt.Printf("B")
	} else if score >=70 {
		grade = "C"
		fmt.Printf("C")
	} else if score >=60 {
		grade = "D"
		fmt.Printf("D")
	} else {
		grade = "E"
		fmt.Printf("E")
	}

	// switch 结构，有限输入类型如星期1-7
	switch grade {
		case "A":
			fmt.Printf("非常优秀")
		case "B":
			fmt.Printf("非常优秀")
		case "C":
			fmt.Printf("非常优秀")
		default:
			fmt.Printf("比较差")
	}

	//循环结构
}
