package main

import "fmt"

//基础加强: 深浅拷贝， 切片传参 ， for range 分析

// 浅拷贝： 仅仅拷贝的是变量的值，没有对指向的空间进行任何的拷贝。go语言中的赋值，函数传参，函数返回值都是浅拷贝
// 深拷贝： 将原有的变量的空间全部拷贝一份

// 切片传参：
/* go中的切片底层是一个结构体
函数传参是浅拷贝
append 当容量不足的时候回动态分配内存，这块内存是我一个新的内存
map 本质是个结构体指针类型，所以把map 作为参数也可以修改元素
*/

// for range 分析
/*
作用： 用于遍历容器类型的数据 例如 数组 切片 map 等
rnage 的本质是一个函数的方法，使用时候可以加括号使用
修改range 得到后的value ，不影响原始切片或者数组
*/
func main() {
	var num int = 10
	ap := &num
	fmt.Println(ap)
	Update(ap)
	fmt.Println(ap)
	fmt.Printf("%d\n", *ap)
	//切片传参  切片底层就是个结构体
	s := make([]int, 5, 5)
	Modify(s)
	fmt.Println(s)

}

// 切片 append

func Update(p *int) {
	*p = 60 // 本质上是对 num 的值修改
}

func Modify(sli []int) {
	for i := 0; i < 5; i++ {
		sli[i] = i
		//sli = append(sli, i) // 追加数据是向新开辟的地址空间追加
	}
	fmt.Println(sli)
	//for range
	for k, v := range sli { // 对形参进行遍历
		sli[k] = 3
		fmt.Println("v:", v)
	}

}
