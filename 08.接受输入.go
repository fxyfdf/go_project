package main

import "fmt"
func main() {
	var age int
	fmt.Println("请输入年龄:")
	//fmt.Scanf("%d",&age)  //将键盘输入赋值给age
	//fmt.Printf("age:%d\n",age)
	//fmt.Printf("age_address:%p",&age) // %p 打印地址空间
	fmt.Scan(&age)
	fmt.Println("age = ",age)
	fmt.Printf("age_address:%p", &age)



}
