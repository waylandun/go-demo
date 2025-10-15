// 教学内容: 函数声明、返回值、多返回值、函数类型、匿名函数与闭包、可变参数。
package main

import "fmt"

// add 提供具名返回值的示例
func add(a, b int) (result int) {
	result = a + b
	return
}

// divide 展示多返回值与错误表示
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// sumAll 使用可变参数
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	total := add(10, 5)
	fmt.Println("具名返回值 add:", total)

	if res, err := divide(10, 2); err == nil {
		fmt.Println("多返回值 divide:", res)
	}

	fmt.Println("可变参数 sumAll:", sumAll(1, 2, 3, 4))

	// 函数作为值
	operate := func(x, y int, op func(int, int) int) int {
		return op(x, y)
	}
	fmt.Println("函数类型:", operate(3, 4, func(a, b int) int {
		return a * b
	}))

	// 闭包: 捕获外部变量
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}

	next := counter()
	fmt.Println("闭包调用:", next(), next(), next())
}
