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

// 排序算法：

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

// 2 选择排序: 每次选择出列表中最大的数然后交换
func SelectSort(list []int) []int {
	// 选择排序，从小到大
	// 1 找出切片中最小的数据  2 和切片中第一个数进行位置交换
	fmt.Println(list)
	fmt.Println(len(list))
	//min := list[0]
	//minIndex :=0
	for i := 0; i < len(list)-1; i++ {
		minIndex := i
		min := list[i]
		for j := i + 1; j < len(list); j++ {
			// 循环找出剩余元素中的最小值和索引位置
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		//如果最小值和剩余元素的第一个索引不相等交换
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
	return list

}

// Map ： 结构，创建 初始化 ， map键值 , map作为函数参数， map案例，
/*
MAP 是一种无序的键值对的集合
map 是通过key 快速检索数据的，key 类似于索引，指向数据的值
键不允许重复
*/
func MapTest() {
	// 数组，切片问题
	names := []string{"str", "lpf", "张三"}
	fmt.Println(names[2])
	//var map1 map[int] string // 键类型 int, 值类型 字符串
	var map1 map[int]string = map[int]string{1: "lfp", 2: "sybs", 33: "张三", 11: "fxyfdf"} // 初始化数据
	//map2 := map[int] string{}
	map2 := map[int]string{1: "lfp", 11: "fxyfdf", 2: "sybs", 33: "张三"} //初始化
	map3 := make(map[string]int)                                        //初始化
	map3["a"] = 1
	map3["ab"] = 2
	fmt.Println("map1 = ", map1, "len(map1) = ", len(map1)) //len 返回的是map 中已有的键值个数
	fmt.Println("map2 = ", map2, "len(map2) = ", len(map2))
	fmt.Println("map3 = ", map3, "len(map3) = ", len(map3))

	// map 键与值
	map4 := map[int]string{1: "lfp", 11: "fxyfdf", 2: "sybs", 33: "张三"}
	fmt.Println("map4[2] = ", map4[2], "len(map4) = ", len(map4))
	// 循环方式获取值
	for k, v := range map4 {
		fmt.Println("k = ", k, "v = ", v)
	}
	// delete(map名字，键)  删除某个值
	fmt.Println(map4)
	delete(map4, 33)
	fmt.Println(map4)
}

// func 函数名（map) {函数体}  调用： 函数名（map)  注意：在函数中修改map 会修改有map 中的值
func MapTest2(map1 map[int]string) {
	// 一个字符串，统计每个字母出现的次数
	for k, v := range map1 {
		fmt.Println("k=", k, "v=", v)
	}
	var str string = "adfereljce adrtrgadd"
	//1 循环每个字母，取出字母
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		ch := str[i] //ch='a' ch='d' ch='f'
		m[ch] = m[ch] + 1
		fmt.Println(m[ch])
	}
	fmt.Println(m)
	//2 统计
	for k, v := range m {
		fmt.Printf("%c:%d \n", k, v)
	}
	//3 输出结果

}

// 结构体
type Student2 struct {
	id   int
	name string
	age  int
	addr string
}

func StructTest(stu Student2) {
	// 结构体： 一系列具有相同类型或不同类型的数据构成的数据集合（结构体可以很好地管理一批有联系的数据，使用结构体可以提高程序的易读性）
	// 结构体可以很好地管理一批有联系的数据，使用结构体可以提高程序的易读性。
	//结构体创建与初始化: 顺序初始化，指定成员初始化，通过结构体变量.成员 完成初始化
	/*
		// 成员前面不加 var
		type Student struct {
			id int
			name string
			age int
			addr string
		}
		// 始化f方式
		var s Student=Student{ id: 101, name:"张三", age:18 ,addr:"BJ"} // 顺序初始化，
		var s1 Student=Student{ id: 101, name:"张三", age:18} // 部分初始化
		var s2 Student  // 通过变量初始化
		s2.id = 3
		s2.name = "lpf"
		s2.addr = "sh"
		fmt.Println(s)
		fmt.Println(s1)
		fmt.Println(s2)

		//结构体与数组：
		// 定义在函数内，结构体只适用于函数内 此处使用 Student2
		var arr[3] Student2=[3] Student2{
			Student2{id: 1,name:"张三",age: 18, addr:"bj"},
			Student2{id: 2,name:"李二",age: 18, addr:"sh"},
			Student2{id: 3,name:"王五",age: 19, addr:"bj"},
		}
		fmt.Println("测试结构体与数组 ")
		fmt.Println(arr)
		fmt.Println(arr[0])
		fmt.Println(arr[0],"age 是 ", arr[0].age)
		arr[0].age = 21
		fmt.Println(arr[0],"age 是 ", arr[0].age)

		// 循环读取
		for i := 0; i< len(arr); i++ {
			//fmt.Println(arr[i])
			fmt.Println(arr[i].name, "age 是", arr[i].age, "岁")
		}
		for k, v := range arr {
			fmt.Printf("%d",k)
			fmt.Println(v.name, "age 是", v.age)
		}

		//结构体与切片: 定义，修改成员值，循环遍历，append 函数使用
		var s3[] Student2 = [] Student2 {
			Student2{id: 1,name:"张三",age: 18, addr:"bj"},
			Student2{id: 2,name:"le",age: 18, addr:"bj"},
			Student2{id: 3,name:"ww",age: 18, addr:"bj"},
		}
		fmt.Println("结构体与切片测试")
		fmt.Println(s3)
		s3[0].age = 21
		fmt.Println(s3)
		// 循环遍历
		for i := 0; i<len(s3); i++{
			fmt.Println(s3[i])
		}
		for _,v := range s3 {
			fmt.Println(v)
		}
		s3 = append(s3,Student2{id:4, name:"李四", age:22, addr:"sh"})
		fmt.Println(s3)

		//结构体与map：  定义 初始化 循环遍历 删除
		m :=make(map[int] Student2)
		m[1]=Student2{id:4333, name:"李四", age:22, addr:"sh"}
		m[2]=Student2{id:53333, name:"李四", age:22, addr:"sh"}
		fmt.Println("测试结构体与map")
		fmt.Println(m)
		delete(m, 2)
		fmt.Println(m)
		fmt.Println(m[1])
		fmt.Println(m[1].name)
		for k,v := range m{
			fmt.Println(k,v)
		}
	*/
	//结构体作为函数  stu
	stu.age = 33
	stu.addr = "函数内修改结构体数据"
	fmt.Println(stu)
	fmt.Println("函数内测试ok")
}

// 注册信息
func InitData(stu []Student2) {
	for i := 0; i < len(stu); i++ {
		fmt.Printf("请输入第%d 学生的详细信息依次为id name age addr\n", i+1)
		fmt.Scan(&stu[i].id, &stu[i].name, &stu[i].age, &stu[i].addr)
	}
	fmt.Println("")
}

// 获取年龄最大
func GetMax(stu []Student2) {
	var max int = stu[0].age
	var maxIndex int
	for i := 0; i < len(stu); i++ {
		if stu[i].age > max {
			max = stu[i].age
			maxIndex = i
		}
	}
	fmt.Println(stu[maxIndex])
}

// 指针
func PointerTest() {

}

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
	//s := []int{11, 3, 4, 2, 5, 33, 22}
	////s2 := BubbleSort(s)
	////fmt.Println(s2)
	//s3 := SelectSort(s)
	//fmt.Println(s3)
	//MapTest()
	//var map2 map[int] string = map[int] string {1:"zs", 2:"ls", 3:"sdr"}
	//MapTest2(map2)
	//stu := Student2{id:5, name:"李四", age:22, addr:"main 函数测试数据"}
	//fmt.Println("主函数开始测试")
	//fmt.Println("1",stu)
	//StructTest(stu)  //函数体内修改结构体数据，不改变结构体外数据
	//fmt.Println("2",stu)
	//MAP
	stu := make([]Student2, 3) // 创建map  3
	InitData(stu)              // 录入学生信息
	GetMax(stu)                //计算最大年龄数谁

}
