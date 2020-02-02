package main

import "User"
import "Product"

// 入口函数
func main() {
	// 运行编译后运行
	Add(1, 2) // 引用同级目录下的文件方法是，编译运行min函数时选择编译当前路径下文件

	//包名.函数名
	User.Login()
	User.GetUser()
	Product.GetProduct()
}
