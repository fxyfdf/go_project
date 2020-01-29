package main

import (
	"fmt"
)

//1数组
/*
var su [3] int //定义数组 st 存储3个数字

*/
func ArrayTest() {
	// 初始化
	// var Numbers[4] int
	var Numbers [5]int = [5]int{2, 3, 1, 3} // 定义数组并初始化， 下标0 开始
	Numbers[4] = 100                        // 部分初始化,未初始化时默认值为0
	Numbers2 := [...]int{7, 8, 111, 9}      // 通过初始化指定数组长度
	fmt.Println(Numbers)
	fmt.Println(Numbers[2])
	fmt.Println(Numbers[4]) // 没有赋值时默认为0
	//fmt.Println(Numbers[7])
	fmt.Println(Numbers2)
	fmt.Println(len(Numbers2))
	var Numbers3 [4]int
	for i := 0; i < len(Numbers3); i++ {
		Numbers3[i] = i + 1
	}
	fmt.Println("list(numbers3) = ", Numbers3)

	// 数组遍历： 1 for ... len()  2 for ... range
	for i := 0; i < len(Numbers3); i++ {
		fmt.Println("Nmubers[", i, "] = ", Numbers3[i])
	}

	//for k,v := range Numbers3{
	for _, v := range Numbers3 { // _ 下划线匿名变量
		fmt.Println(v)
	}

}

func getPrint(n [5]int) {
	// 数组作为函数参数
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
}

func compareValue(n1 [5]int, n2 [5]int) bool {
	// 比较两个数组中的元素是否一致 1 长度，2 对应的值
	var b bool = true
	// 1 对比两个数组长度是否一致
	if len(n1) == len(n2) {
		// 2 判断对应的值
		for i := 0; i < len(n1); i++ {
			if n1[i] == n2[i] {
				continue
			} else {
				b = false
				fmt.Println("数组中，值不一致")
				break
			}
		}
	} else {
		b = false
		fmt.Println("数组长度不一致")
	}
	return b
}

func ArrayTest2() {
	// 二维数组定义，初始化与赋值， 循环遍历打印
	var arr2 [2][3]int //定义一个两行三列的
	arr2[0][1] = 12    // 部分初始化
	arr2[1][1] = 11
	var arr21 [2][3]int = [2][3]int{{1, 2, 3}, {21, 22, 23}} // 全部初始化
	var arr22 [2][3]int = [2][3]int{{1, 2}, {23}}            // 部分初始化
	var arr23 [2][3]int = [2][3]int{1: {5, 6}}               //指定元素初始化，第二行1，2列为5，6
	arr24 := [...][3]int{{1, 2, 3}, {5, 6}}                  //行可以用... ,列不可以用代替
	fmt.Println(arr2)
	fmt.Println(arr21)
	fmt.Println(arr22)
	fmt.Println(arr23)
	fmt.Println(arr24)
	// 循环方式打印
	fmt.Println(len(arr2))    //输出有几行
	fmt.Println(len(arr2[0])) //输出几列
	fmt.Println(arr2)
	for i := 0; i < len(arr2); i++ { // 循环行
		for j := 0; j < len((arr2[0])); j++ { //循环列
			fmt.Println(arr2[i][j])
		}
	}
	for i, v := range arr2 {
		fmt.Println("i", i)
		fmt.Println("v", v) // 第几行的数组
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

}

// 切片 --> 可以理解成动态的数组，但不是数组
func SliceUp1() {
	// 数组的问题，长度固定，改变长度时需要重新定义
	//var num[5] int = [5] int {1,2,3,4,5} //需要改变 存储个数时，需要重新定义
	// 切片创建
	var sliceUp1 []int  // 默认创建的切片长度为0
	sliceUp2 := []int{} //大括号写初始化值
	//sliceUp21 := []int{1, 3, 3, 3}  //大括号写初始化值
	sliceUp3 := make([]int, 3, 5)           //参数： 切片类型，长度（已经初始化的空间），容量(已经开辟的空间：包括初始化和未初始化的 。 容量不写时， 容量等于长度
	sliceUp1 = append(sliceUp1, 1, 2, 3, 4) // 切片追加 元素
	sliceUp1[3] = 4                         // 修改已有值，注意下标越界问题
	sliceUp3[2] = 4                         // 修改已有值，注意下标越界问题
	fmt.Println(sliceUp1, "len(sliceUp1) = ", len(sliceUp1))
	fmt.Println("sliceUp1[3] = ", sliceUp1[3]) // 通过下标获取 值
	fmt.Println(sliceUp2, "len(sliceUp2) = ", len(sliceUp2))
	fmt.Println(sliceUp3, "len(sliceUp3) = ", len(sliceUp3)) // len() 返回的时长度
	fmt.Println(sliceUp3, "cap(sliceUp3) = ", cap(sliceUp3)) //cap() 返回的时容量

	// 循环变量
	for i := 0; i < len(sliceUp3); i++ {
		sliceUp3[i] = i + 10
	}
	fmt.Println(sliceUp3, "len(sliceUp3) = ", len(sliceUp3)) // len() 返回的时长度
	sliceUp3 = append(sliceUp3, 1, 2, 3, 4, 4)
	fmt.Println(sliceUp3, "len(sliceUp3) = ", len(sliceUp3)) // len() 返回的时长度
	sliceUp5 := []int{1, 2, 3, 4, 5, 5}
	for i, v := range sliceUp5 {
		fmt.Println("i = ", i, "v = ", v)
	}

	// 截取部分元素
	fmt.Println(sliceUp5)
	sliceUp6 := sliceUp5[1:3:5] //0 开始 3 结束不包含下标为3, 容量 5(切片中最多容纳多少个元素， 容量=第三个值-第一个值， 长度=第二个值-第一个值)
	sliceUp7 := sliceUp5[:]     // [:] [1:] [:1]  [1:2:5]  len()长度  cap()容量
	fmt.Println(sliceUp6)
	fmt.Println(sliceUp7)
	fmt.Println(cap(sliceUp6))
	fmt.Println(len(sliceUp6))
	// 切片值修改
	sliceUp9 := sliceUp5
	sliceUp8 := sliceUp5[2:5]
	sliceUp8[0] = 88 //切片修改值后，原有切片值会修改，截取的切片数据本质上时 原来切片的索引并没有重新开辟地址空间

	fmt.Println(sliceUp8)
	fmt.Println(sliceUp5)
	fmt.Println(sliceUp9)
	// append 追加元素， 容量超过后扩容 容量*2 超过1024 时，扩容为原有的1/4
	sliceUp10 := make([]int, 3, 5) // 原有容量为5，第一次扩容后为10
	sliceUp10 = append(sliceUp10, 1, 2, 3, 4, 5, 6)
	fmt.Println(sliceUp10, len(sliceUp3), cap(sliceUp10))
	// copy 拷贝  copy(s1, s2)  将s2 中的切片拷贝到 s1 中  只同步位置相同的内容
	sliceUp11 := []int{11, 111}
	sliceUp12 := []int{2, 22, 23}
	copy(sliceUp12, sliceUp11)
	fmt.Println(sliceUp11)
	fmt.Println(sliceUp12)

	// 计算一组整数之和，1 确定数据个数 2 输入求和数据，并存储到切片中
	//fmt.Println("请输入求和数据个个数：")
	//var count int
	//fmt.Scan(&count)
	//s := make([] int,)
}

func SliceUp2(num []int) {
	//切片作为函数参数
	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}
}

func SliceUp3() {
	//输出一组整数中最大的值
	// 1 比较数据个数
	fmt.Println("请输入数据个数")
	var count int
	fmt.Scan(&count)
	// 2 请输入比较的数，并且存放在切片中
	s := make([]int, count) //容量大小为个数
	AddData(s)
	// 3 进行比较,返回最大值
	MaxS := MaxValue(s)
	fmt.Println("max = ", MaxS)
}

func AddData(num []int) {
	// 存储输入数据
	for i := 0; i < len(num); i++ {
		fmt.Println("请输入第%d 个数", i)
		fmt.Scan(&num[i]) //修改截取切片数据会改变原始数据
	}
}

func MaxValue(num []int) int {
	var max int = num[0]
	for i := 0; i < len(num); i++ {
		if num[i] > max {
			max = num[i]
		}
	}
	return max
}

// 排序算法

// 1 冒泡排序
func BubbleSort(list []int) []int {
	//从小到大排序
	// 1 排序找出最大数 循环，每次减1
	//var i int //一共循环次数
	//var j int //每次比较次数
	fmt.Println(list)
	fmt.Println(len(list))
	var temp int
	// 把最大的值冒泡上去，
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ { //  1-n,0-n-1,0-n-1
			if list[j] > list[j+1] {
				temp = list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}
	return list
}

// 2 选择排序

// Map

// 结构体

// 指针

// 其他

func main() {
	//var Numbers[5] int = [5]int{1, 2, 3, 4, 5}
	//var Numbers2[5] int = [5]int{1, 2, 33, 4, 5}
	// ArrayTest()
	// getPrint(Numbers)
	//compareValue(Numbers,Numbers2)
	//fmt.Println(compareValue(Numbers,Numbers2))
	//fmt.Println(Numbers2 == Numbers)
	//ArrayTest2()
	//SliceUp1()
	//s := make([] int,10)
	//s = append(s,1,2,3,4,5,6,7)
	//SliceUp2(s)
	//SliceUp3()
	s := []int{1, 3, 4, 2, 5, 33, 22}
	s2 := BubbleSort(s)
	fmt.Println(s2)
}
