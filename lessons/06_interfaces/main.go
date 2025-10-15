// 教学内容: 接口定义、隐式实现、接口组合、类型断言与类型选择、空接口。
package main

import "fmt"

// Speaker 定义一个行为契约
type Speaker interface {
	Speak() string
}

// Runner 演示接口组合
type Runner interface {
	Speaker
	Run() string
}

type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return "我是机器人 " + r.Model
}

func (r Robot) Run() string {
	return "机器人正在奔跑"
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "你好，我是 " + p.Name
}

func describeSpeaker(s Speaker) {
	fmt.Println("接口调用 Speak:", s.Speak())

	// 类型断言: 判断具体类型
	if robot, ok := s.(Robot); ok {
		fmt.Println("通过类型断言识别 Robot:", robot.Model)
	}
}

func main() {
	var s Speaker

	s = Person{Name: "Bob"}
	describeSpeaker(s)

	s = Robot{Model: "RX100"}
	describeSpeaker(s)

	// 接口组合示例
	var runner Runner = Robot{Model: "RX200"}
	fmt.Println("runner Speak:", runner.Speak())
	fmt.Println("runner Run:", runner.Run())

	// 类型选择 type switch
	printType := func(v interface{}) {
		switch value := v.(type) {
		case string:
			fmt.Println("字符串:", value)
		case int:
			fmt.Println("整型:", value)
		case Speaker:
			fmt.Println("实现 Speaker:", value.Speak())
		default:
			fmt.Println("未知类型")
		}
	}

	printType("hello")
	printType(42)
	printType(Person{Name: "Carol"})
}
