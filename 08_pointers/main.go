// 教学内容: 指针基础、取地址与解引用、函数指针参数、new 与 make 的区别、结构体指针。
package main

import "fmt"

func incrementByPointer(n *int) {
	*n++
}

type Node struct {
	Value int
	Next  *Node
}

func main() {
	value := 10
	ptr := &value

	fmt.Printf("指针地址: %p 指向的值: %d\n", ptr, *ptr)

	incrementByPointer(ptr)
	fmt.Println("通过指针修改后的值:", value)

	// new 分配零值并返回指针
	numPtr := new(int)
	fmt.Println("new 返回的指针零值:", *numPtr)
	*numPtr = 5
	fmt.Println("修改 new 指针的值:", *numPtr)

	// make 只能用于切片、map、chan，用于初始化底层结构
	slice := make([]int, 0, 3)
	fmt.Printf("make 创建切片: %v len=%d cap=%d\n", slice, len(slice), cap(slice))

	// 结构体指针常与链表等数据结构结合
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	fmt.Printf("链表: %d -> %d\n", head.Value, head.Next.Value)
}
