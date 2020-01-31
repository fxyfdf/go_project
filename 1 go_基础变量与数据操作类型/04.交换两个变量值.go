package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int = 20
	/*	var temp int
		temp = num1
		num1 = num2
		num2 = temp*/
	num1, num2 = num2, num1
	fmt.Println(num1, num2)
}
