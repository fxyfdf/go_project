package main

import "fmt"

/*
1 函数定义
2 函数调用
3 函数参数
通过下标方式获取参数值
通过循环的方式获取参数值
固定参数放在前面，不定参数放在后面
在对函数进场调用时，固定参数必须传值，不定参数可以根据需要决定
func 函数名 (参数){
	方法
}

递归函数，递归函数基本语法

*/
func PlayGame() {
	fmt.Println("超级玛丽 飞一飞")
	fmt.Println("超级玛丽 飞一飞")
}

func Play() {
	fmt.Println("屏幕闪烁")
	fmt.Println("播放音乐")
}

// 参数函数
func SumAdd(num int) {
	var sum int
	for i := 1; i <= num; i++ {
		sum += i
	}
	fmt.Println("1-", num, "之和为", sum)
}

// 传入两个参数， 形参和实参个数和类型要一致
func Add(num int, str string) {
	fmt.Println("传入整数num 为：", num)
	fmt.Println("传入整数str 为：", str)
}

// 函数不定参数列表
func TestSum(args ...int) {
	var sum int = 0
	for i := 0; i < len(args); i++ {
		sum = sum + args[i]
	}
	fmt.Println(args, "sum = ", sum)
	// 集合操作
	for k, v := range args {
		// for _, v := range args { //_ 匿名变量
		fmt.Println("k = ", k)
		fmt.Println("v = ", v) // 存储的具体的值
	}

}

// 函数返回值
//
func AddResult(num1 int, num2 int) (sum int) {
	//int表示指定函数返回的数据类型  {sum int}
	//var sum int
	sum = num1 + num2
	return sum // 将变量sum中存储的值返回
	//return   //写{sum int} 时不需要指定返回值
}

func GetResult() (int, int) {
	var num1 int = 10
	var num2 int = 20
	return num1, num2
}

// 函数案例  函数作用域  延迟调用  函数应用总结
// 模拟用户注册： 输入 用户名 密码 邮箱
// 三个函数
// 用户注册函数
func Register() string {
	// 1 接收用户输入信息
	var userName string = "admin"
	var userPwd string = "123123"
	var userEmail string = "aa@qq.com"
	var b string
	// 2 校验校验
	b = CheckInfo(userName, userPwd, userEmail)
	fmt.Println(b)
	// 3 校验成功后发送邮件
	if b == "true" {
		fmt.Println(SendEmail())
	} else {
		return "信息有误"
	}
	return "结束"
}

// 校验信息
func CheckInfo(name string, pwd string, email string) (b string) {
	if name != "" && pwd != "" && email != "" {
		b = "true"
		return b
	} else {
		b = "false"
		return b
	}
}

// 发送邮件
func SendEmail() string {
	return "邮件发送成功"
}

// 函数作用域
var b int = 10 //函数外的变量为全局变量
func TestA() {
	a := 1 // 此变量为局部变量，其他函数不能使用
	fmt.Println("b = ", b)
	b = 20 // 函数内修改全局变量，执行此函数后变量会变为20 ，其他函数在使用时是20
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
func TestB() {
	fmt.Println("b = ", b)
}

// 延迟调用  defer 执行顺序，后进先出  如关闭文件
func DeferTest() {
	defer fmt.Println("第一行") //先defer  最后执行
	fmt.Println("第二行")
	defer fmt.Println("第三行")
	fmt.Println("第四行")

}

// 递归函数：  递归函数基本语法 递归解决树形结构问题，   如果一个函数在内部调用自身本身，这个函数就是递归函数

//自己调用自己，但是一定要有一个出口
func Test(n int) int {
	// 只有第一排知道排数
	if n == 1 {
		return 1
	}
	// 如果部署第一排问前一排
	r := Test(n - 1)
	fmt.Println("前一排的排数 ：", r)
	// 把前一排加一，计算自己的排数
	return r + 1
}

// 计数一个数的阶乘
var s int = 1 // 用于存储阶乘值
func TestDemo(n int) {
	if n == 1 {
		return // 终止跳出函数
	}
	s *= n
	TestDemo(n - 1)
}

func main() {
	// 函数调用
	//PlayGame()
	//Play()
	//SumAdd(10)
	//Add(100, "hello world")
	//TestSum(1,2,3,4,5)
	//var s int
	//s = AddResult(12, 12)
	//fmt.Println("num1 + num2 = ", s)
	//fmt.Println(GetResult())
	//Register()
	//TestA()
	//TestB()
	//DeferTest()
	c := Test(3)
	fmt.Println(c)
	TestDemo(10)
	fmt.Println(s)
}
