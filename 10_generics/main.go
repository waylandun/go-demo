// 教学内容: 泛型函数、类型参数约束、泛型类型、comparable 约束与自定义接口约束。
package main

import "fmt"

// Add 支持任意 number 类型，通过 ~ 运算符合并底层类型
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// Pair 泛型结构体示例
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// MapKeys 提供约束为可比较类型的键
func MapKeys[K comparable, V any](pairs []Pair[K, V]) map[K]V {
	result := make(map[K]V)
	for _, p := range pairs {
		result[p.Key] = p.Value
	}
	return result
}

func main() {
	ints := []int{1, 2, 3}
	floats := []float64{1.5, 2.5, 3.0}
	fmt.Println("Sum[int]:", Sum(ints))
	fmt.Println("Sum[float64]:", Sum(floats))

	pairs := []Pair[string, int]{
		{Key: "Alice", Value: 90},
		{Key: "Bob", Value: 85},
	}

	fmt.Println("泛型 map:", MapKeys(pairs))
}
