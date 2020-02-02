package User

import "fmt"

func Login() {
	fmt.Println("用户登录")
}

// 如果想要在其他文件内调用该方法，该方法名应该大写首字母
func GetUser() {
	fmt.Println("获取用户信息")
}
