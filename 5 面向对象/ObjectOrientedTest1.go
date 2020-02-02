package main

import "fmt"

/*
练习：
定义一个学习类，有6个属性，分别为姓名，性别，年龄，语文，数学，英语成绩
第一个方法： 打招呼方法： 介绍自己叫**，今年**岁，性别 *
第二个方法： 计算总分与平均分
*/
// 第一个学生类
type Student struct {
	name    string
	sex     string
	age     int
	chinese int
	math    int
	english int
}

// 2 添加方法
func (s Student) SayHello(username string, userAge int, userSex string) {
	//2.1 初始化
	s.name = username
	s.age = userAge
	s.sex = userSex
	// 2.2 初始化判断
	if s.sex != "男" && s.sex != "女" {
		s.sex = "男"
	}
	if s.age < 1 || s.age > 100 {
		s.age = 18
	}
	// 2.3 打印输出
	fmt.Printf("我叫%s,年龄 %d, 性别 %s", s.name, s.age, s.sex)

}

func main() {
	var stu Student
	stu.SayHello("zs", 17, "")
}
