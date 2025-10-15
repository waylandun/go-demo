// 教学内容: 数组、切片、map 的创建、操作、容量管理以及内建函数 make/len/cap/append/copy。
package main

import "fmt"

func main() {
	// 数组: 固定长度，值类型
	var scores [3]int = [3]int{95, 88, 76}
	fmt.Println("数组:", scores)

	// 切片: 基于数组的引用类型，动态长度
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5)
	fmt.Printf("切片: data=%v len=%d cap=%d\n", numbers, len(numbers), cap(numbers))

	// 使用 make 创建切片并扩容
	buffer := make([]int, 2, 4)
	buffer[0], buffer[1] = 10, 20
	buffer = append(buffer, 30)
	fmt.Printf("make 切片: data=%v len=%d cap=%d\n", buffer, len(buffer), cap(buffer))

	// copy 将一个切片复制到另一个切片
	dest := make([]int, len(numbers))
	copy(dest, numbers)
	fmt.Println("copy 后:", dest)

	// map: key-value 存储，使用 make 初始化
	userAges := make(map[string]int)
	userAges["Alice"] = 25
	userAges["Bob"] = 30
	fmt.Println("map 访问:", userAges["Alice"])

	// 判断键是否存在
	if age, ok := userAges["Charlie"]; ok {
		fmt.Println("找到 Charlie:", age)
	} else {
		fmt.Println("Charlie 不在 map 中")
	}

	// 遍历 map 注意无序
	for name, age := range userAges {
		fmt.Printf("map 遍历: %s -> %d\n", name, age)
	}

	// 删除键
	delete(userAges, "Bob")
	fmt.Println("delete 后:", userAges)
}
