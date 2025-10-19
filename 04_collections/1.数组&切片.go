package main

import (
	"fmt"
	"strings"
)

// go 中如果是数组，默认是值传递
//func printArry(arr [5]int){}

// 切片 - 才是引用传递 和Python的列表一样
func printArry(arr []int) {
	// 更改切片内容 外部会同步更改 因为是指针（引用传递）
	arr[0] = 886
}

func main() {
	// 固定长度数组 - 设定了固定长度的叫数组，反之是切片
	var mlist [5]int
	// 按索引位置赋值
	mlist[0] = 1
	mlist[1] = 2
	mlist[2] = 3
	fmt.Println("数组1:", mlist)
	// 2.定义的同时直接赋值
	mlist2 := [5]int{1, 2, 3}
	fmt.Println("数组2:", mlist2)
	// 3. 不限制长度的数组
	mlist3 := [...]int{}
	fmt.Println("数组3:", mlist3)

	// 😂 固定长度数组不能直接使用append，需要先转换为切片
	// append 追加元素到切片
	mlistSlice := mlist[:]
	mlistSlice = append(mlistSlice, 4)

	// 导包 strings 注意和 string 区分
	// 字符串重复 类似 Python中的 print("---"*30)
	fmt.Println(strings.Repeat("---", 30))

	// 切片 - 中括号中不写任何内容
	// 定义1：(❌不推荐)声明了一个空切片，但没有分配空间 -
	var q1 []int
	q1 = append(q1, 1)
	// 需要通过make函数分配空间 注意这里是 = 而不是 :=
	//q1 = make([]int, 5)
	fmt.Println("切片1:", q1)
	// 定义2: 声明并初始化(空，不会有任何元素 和Python中 a = [] 一致)
	q2 := []int{}
	fmt.Println("切片2:", q2)
	// 定义2的延伸，带初始值
	q2_1 := []int{1, 2, 3, 4, 5}
	fmt.Println("切片2延伸:", q2_1)

	//// 定义4 使用make 定义切片 初始长度为5 可动态扩容 2倍扩充
	//q4 := make([]int, 5)
	//fmt.Println("切片4:", q4)
	//
	//for i := 0; i < 7; i++ {
	//	q1 = append(q1, i)
	//}
	//fmt.Println("切片1:", q1)
	//
	//// 遍历切片1
	//for i := 0; i < len(q1); i++ {
	//	fmt.Println("方法1-遍历切片:", q1[i])
	//}
	//// range 遍历会返回两个对象，一个是索引，一个是值
	////for index, value := range q1 {
	//for _, v := range q1 {
	//	fmt.Println("方法2-range遍历切片:", v)
	//}
	//fmt.Println("---------------------------------------------")
	//printArry(q1)
	//fmt.Println("切片1 第0个元素被改变:", q1)

}
