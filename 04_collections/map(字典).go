package main

import "fmt"

func main() {
	// 定义一个map
	//var userAges map[string]any  // 这只是声明 没有初始化 不能用
	// 未初始化时赋值会报错
	//userAges["name"] = "lisi"
	//userAges["age"] = 15

	// 定义一个map 的常用方法
	userInfo := make(map[string]any)
	userInfo["name"] = "lisi"
	userInfo["age"] = 15
	fmt.Println("map:", userInfo)
}
