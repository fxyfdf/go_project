# Go5期day01课上笔记



# 课程物料

## 预习URL

[讲师博客](https://www.liwenzhou.com/posts/Go/go_menu/)

## 课上笔记

markdown格式的笔记+截图

## 视频

课后自己拷或者等百度网盘连接

## 课后作业

回去之后自己注册[路飞学城](www.luffycity.com)账号，把自己的用户名、真实姓名告诉班主任

每周作业要求都在路飞学城上会有展示，每周五12:00之前把作业打包提交到路飞学城。



![1561777113169](D:\Go5期\assets\1561777113169.png)

# Go语言开发环境搭建

## 安装Go开发包

### 下载

[https://golang.google.cn/dl/](https://golang.google.cn/dl/)

![1561777435197](D:\Go5期\assets\1561777435197.png)

安装目录一定选一个好记的。

![1550236972659](https://www.liwenzhou.com/images/Go/install_go_dev/1550236972659.png)

安装完成后，输入`go version`查看go版本号。

![1561777659521](D:\Go5期\assets\1561777659521.png)

## 配置GOPATH

![1561778001060](D:\Go5期\assets\1561778001060.png)

![1561778199604](D:\Go5期\assets\1561778199604.png)

详细步骤：

1. 在自己的电脑上新建一个目录`D:\go`（存放我编写的Go语言代码）
2. 在环境变量里，新建一项：`GOPATH:D:\go`
3. 在`D:\go`下新建三个文件夹，分别是：`bin`、`src`、`pkg`
4. 把`D:\go\bin`这个目录添加到`PATH`这个环境变量的后面
   1. Win7是英文的`;`分隔
   2. Win10是单独一行
5. 你电脑上`GOPATH`应该是有默认值的，通常是`%USERPROFILE%/go`， 你把这一项删掉，自己按照上面的步骤新建一个就可以了。 

![1561780163679](D:\Go5期\assets\1561780163679.png)

## GO语言项目结构

![1561780567275](D:\Go5期\assets\1561780567275.png)

## 下载并安装VS Code



### 下载VSCODE

[官方下载连接](https://code.visualstudio.com/Download)

### 安装

“下一步安装法”

### 安装中文插件包

![1561780951477](D:\Go5期\assets\1561780951477.png)

### 安装Go扩展

![1561781081386](D:\Go5期\assets\1561781081386.png)

### VSCode界面介绍

![1550240342443](https://www.liwenzhou.com/images/Go/install_go_dev/1550240342443.png)

## 编译go build

使用`go build`

1. 在项目目录下执行`go build`
2. 在其他路径下执行`go build`， 需要在后面加上项目的路径（项目路径从GOPATH/src后开始写起，编译之后的可执行文件就保存在当前目录下）
3. `go build -o hello.exe`



### go run

像执行脚本文件一样执行Go代码



### go install

`go install`分为两步：

	1. 先编译得到一个可执行文件
 	2. 将可执行文件拷贝到`GOPATH/bin`

### 交叉编译

Go支持跨平台编译

例如：在windows平台编译一个能在linux平台执行的可执行文件

```bash
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

执行`go build`

Mac平台交叉编译：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

# Go语言文件的基本结构

```go
import "fmt"

// 函数外只能防止标识符（变量\常量\函数\类型）的声明
// fmt.Println("Hello") // 非法

// 程序的入口函数
func main() {
	fmt.Println("Hello world!")
}

```

# 变量和常量

[博客](https://www.liwenzhou.com/posts/Go/01_var_and_const/)

## 变量

Go语言中的变量必须先声明再使用

### 声明变量

`var s1 string`:声明一个保存字符串类型数据的s1变量

### 注意事项

变量名推荐使用驼峰式

变量必须先声明再使用

变量声明且赋值后必须使用

# 基本数据类型

## 整型

## 浮点型

## 复数

## 布尔值

## 字符串

Go语言中字符串是用双引号包裹的！！！

Go语言中单引号包裹的是字符！！！

```go
// 字符串
s := "Hello 沙河"
// 单独的字母、汉字、符号表示一个字符
c1 := 'h'
c2 := '1'
c3 := '沙'
// 字节： 1字节=8Bit(8个二进制位)
// 1个字符‘A’=1个字节
// 1个utf8编码的汉字'沙'= 一般占3个字节
```

## byte和rune

# 流程控制

## if判断

```go
// if条件判断
func main() {
	// age := 19

	// if age > 18 { // 如果 age > 18 就执行这个{}中的代码
	// 	fmt.Println("澳门首家线上赌场开业啦！")
	// } else { // 否则就执行这个{}中的代码
	// 	fmt.Println("改写暑假作业啦！")
	// }

	// 多个判断条件
	// if age > 35 {
	// 	fmt.Println("人到中年")
	// } else if age > 18 {
	// 	fmt.Println("青年")
	// } else {
	// 	fmt.Println("好好学习！")
	// }

	// 作用域
	// age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 { // 如果 age > 18 就执行这个{}中的代码
		fmt.Println("澳门首家线上赌场开业啦！")
	} else { // 否则就执行这个{}中的代码
		fmt.Println("改写暑假作业啦！")
	}

	// fmt.Println(age) // 在这里是找不到age
}
```

## for循环

```go
func main() {
	// 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种1
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// 变种2
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 无限循环
	// for {
	// 	fmt.Println("123")
	// }

	// for range循环
	s := "Hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
```





# 作业

1. 常量声明及练习题回去自己写一下，特别是iota定义数量级时（1024）
2. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用`fmt.Printf()`搭配`%T`分别打印出上述变量的值和类型。
3. 编写代码统计出字符串`"hello沙河小王子"`中汉字的数量。(自己查资料)
4. ![1561805655415](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\1561805655415.png)