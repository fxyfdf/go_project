package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

//1 异常处理 2 文本文件处理 3 字符串处理 4 工程管理

//1 异常处理:
func ExceptionTest(num1 int, num2 int) (result int, err error) { //等于创建了两个变量： result err
	//异常简介:
	//result := num1 / num2

	//error接口: 异常的处理

	// 处理 panic: runtime error: integer divide by zero  除数为0 时的异常
	err = nil
	if num2 == 0 {
		err = errors.New("除数不能为0") //构建一个异常信息返回
		return                     // 返回异常信息终止执行
	}
	result = num1 / num2
	return

	//panic 函数，
	/*
		//var num[10] int
		//num[num1] = 12  // 测试数组下标越界异常：panic: runtime error: index out of range
		fmt.Println("panic 函数测试")
		// 程序员本身不会调用该函数，但是如果程序中出现比较严重的异常，系统会自动调用该函数终止程序执行
		panic("abc")  // 传个字符串， 引发异常强制终止整个程序的执行
		fmt.Println("hello")

		//return result
	*/

	//异常处理总结
}

//recover 错误拦截器，
func TestRecover1(n int) {
	defer TestRecover2() // 拦截异常函数  recover 只有在defer 调用的函数中有效
	var num [10]int
	num[n] = 11
	fmt.Println(num)
	fmt.Println("aaa")

}
func TestRecover2() { // recover 只有在defer 调用的函数中有效
	//recover()  // 拦截所有异常
	fmt.Println(recover()) // 输出异常信息
}

//2 文本文件处理
func FileManage1() {
	// 创建文件，写入数据，读取文件数据，文本文件案例，文件处理总结
	// 创建文件： 导入os包， 指定文件路径名， create() 创建文件后记得关闭文件
	// 创建文件需要指定文件的存储路径以及文件名称 注意 / d:/test/file.txt  创建文件时文件路径需要存在
	file, err := os.Create("file.txt") // 返回文件指针和错误信息
	// 判断是否有异常，不论是否有异常都需要关闭文件
	if err != nil {
		fmt.Println(err)
		//file.Close()
	}
	defer file.Close() //延迟最后执行
	//关闭
	//file.Close()
	//写入数据: WriteString Write WriteAt
	n, err := file.WriteString("我是字符串信息一行的\n") // n字符串长度，err错误信息,重写
	var str2 string = "Hello World"
	n2, err2 := file.Write([]byte(str2)) // 需要将字符串转换成字节切片, 本质上和WriteString 一样
	var str3 string = "00003 WritrAt0000"
	n3, err3 := file.WriteAt([]byte(str3), 0) //在指定位置写入数据
	num, _ := file.Seek(0, io.SeekEnd)        //将光标定位到原有内容的后面  0，原有文件后面加0个空格，10 加10个空格
	file.WriteAt([]byte(str3), num)           //在获取到的最后光标位置num 后面写入数据
	file.WriteAt([]byte(str3), 100)           //在获取到的最后光标位置num 后面写入数据
	fmt.Println("num=", num)                  //原有数据长度
	if err != nil {
		fmt.Println("错误信息为：", err)
	}
	fmt.Println("写入字符串长度为：", n)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("写入字节长度为", n2)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println("写入字节长度为", n3)
	// 向已存在的文件中写入文件  打开文件， OpenFile()

	//os.Create("file2.txt")
	//                               打开文件   打开模式追加     文件权限
	file1, fileerr1 := os.OpenFile("file2.txt", os.O_APPEND, 6)
	if fileerr1 != nil {
		fmt.Println(fileerr1)
	}
	defer file1.Close()
	// 通过文件指针向文件中写入数据，读取数据
	file2, fileerr2 := file1.WriteString("我是测试数据1")
	if fileerr2 != nil {
		fmt.Println(fileerr2)
	}
	fmt.Println(file2)

	//读取文件数据:  基本流程  打开要读取的文件， 对文件进行读取， 关闭文件

	// 1 打开文件:  open == OpenFile(name, O_RDONLY, 0)  ，以只读的方式打开
	file3, fileerr3 := os.Open("file2.txt")
	if fileerr3 != nil {
		fmt.Println(fileerr3)
	}
	defer file3.Close()
	//2 文件读取: 定义一个字符类型切片，存储从文件中读取的数据
	//buffer := make([]byte,1024*2) // 大小为2kb  如果文件太大可以循环打印
	buffer := make([]byte, 10) // 读取9个字符
	//n3, n3err3 := file3.Read(buffer)  //把读取出来的数据存储到buffer 里面
	for { //进行死循环
		n3, n3err3 := file3.Read(buffer) //把读取出来的数据存储到buffer 里面
		if n3err3 == io.EOF {            // 如果到达文件末尾结束
			break
		}
		//fmt.Println("读取数据:",buffer[:n3])
		//fmt.Print("读取数据:",string(buffer[:n3]))
		fmt.Print(string(buffer[:n3]))
		//fmt.Println(string(buffer[:n3]))
	}
	//if n3err3 != nil {
	//	fmt.Println(n3err3)
	//}
	//fmt.Println("读取数据长度",n3)
	//fmt.Println("读取数据:",buffer)
	//fmt.Println("读取数据:",buffer[:n3])
	//fmt.Println("读取数据:",string(buffer[:n3]))
	//3 关闭文件

	// 循环读取文件：

	//文本文件案例，文件处理总结
}

// 文本操作案例， 模拟文件备份
func FileBak() {
	//1 读取文件
	file1, err1 := os.Open("file2.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	//2 创建新文件
	file2, err2 := os.Create("file3.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer file1.Close()
	defer file2.Close()
	//3 将源文件file1.txt读取出来，　死循环　写入新文件
	buffer := make([]byte, 1024) // 创建一个字符类型切片
	for {
		// 读取长度，异常信息
		n1, n1err := file1.Read(buffer) // 把读取处理的内容存放到buffer 内
		// 如果到文件末尾结束
		if n1err == io.EOF {
			break
		}
		// 把每次读取处理的内容 写入file2
		file2.Write(buffer[:n1])
	}
}

// 3 字符串处理
func StringOperation() {
	//https://studygolang.com/pkgdoc
	// 1 Contains(s, substr string) bool  判断字符串s是否包含子串substr。
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	// 2 Join(a []string, sep string) string 将一系列字符串连接为一个字符串，之间用sep来分隔。
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // 以逗号对切片中内容进行连接
	// 3 index 	 Index(s, sep string) int 子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	fmt.Println(strings.Index("chicken", "ken")) // 位置从0 开始
	fmt.Println(strings.Index("chicken", "dmr"))
	// 4 Repeat Repeat(s string, count int) string 返回count个s串联的字符串。
	fmt.Println("ba" + strings.Repeat("na", 2))
	// 5 replace 替换字符串  返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	// 6 Split
	str6 := "hello:world:lll"
	s6 := strings.Split(str6, ":")
	fmt.Println(s6)
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

}

// 工程管理:  将不同的代码写到不同的文件中
// 创建同级目录： 所有的文件放置同一个文件夹下  如： 创建src目录，在该目录下创建go源码文件（注意创建的源码文件在同一个包下)， 在进行配置build目录运行

//创建不同级目录 :  将代码文件放在不同的文件内进行管理
// 在src 目录下创建不同目录
// 进行配置 : 与同级目录时配置相同
// 函数调用： 函数名称的首字母要大写，要带入相应的包

//工程管理总结

func PrijectManage() {

}

func main() {
	// 构建异常信息展示
	//num,err := ExceptionTest(12, 0)
	//if err!=nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(num)
	//
	//}
	//revover
	//TestRecover1(11)
	//os.Open()  //Ctrl + b 快速查看方法
	// 文本文件处理
	//FileManage1()
	//FileBak()
	StringOperation()
}
