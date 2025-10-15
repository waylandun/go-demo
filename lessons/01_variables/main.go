// 教学内容: 变量与常量的声明、类型推断、零值、短变量声明以及批量声明。
package main

import "fmt"

func main() {
	// 零值: 声明但未初始化的变量会得到类型的零值
	var zeroInt int
	var zeroString string
	fmt.Println("零值示例:", zeroInt, zeroString)

	// 显式类型声明与初始化
	var age int = 18
	var name string = "Gopher"
	fmt.Printf("显式声明: age=%d name=%s\n", age, name)

	// 类型推断: 可以省略类型，由编译器推断
	var temperature = 36.5
	fmt.Printf("类型推断: temperature=%v (类型: %T)\n", temperature, temperature)

	// 短变量声明: 只能在函数内部使用的 := 语法
	height := 172
	fmt.Printf("短变量声明: height=%d\n", height)

	// 批量声明: 同时声明多个变量或常量
	var (
		width, length = 5.5, 10.2
		shape         = "rectangle"
	)
	const (
		version  = "v1.0.0"
		isActive = true
	)

	fmt.Printf("批量变量: width=%.1f length=%.1f shape=%s\n", width, length, shape)
	fmt.Printf("批量常量: version=%s isActive=%t\n", version, isActive)
}
