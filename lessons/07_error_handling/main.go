// 教学内容: 错误处理模式、errors 包、自定义错误类型、panic 与 recover、安全资源释放。
package main

import (
	"errors"
	"fmt"
)

// DivideError 自定义错误类型
type DivideError struct {
	Dividend float64
}

func (e DivideError) Error() string {
	return fmt.Sprintf("不能使用 %v 作为除数", e.Dividend)
}

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivideError{Dividend: b}
	}
	return a / b, nil
}

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover 捕获 panic:", r)
		}
	}()

	fmt.Println("开始执行风险操作...")
	panic("发生严重错误，触发 panic")
}

func main() {
	// 使用 errors.New
	err := errors.New("这是一个简单错误")
	fmt.Println("errors.New:", err)

	// fmt.Errorf 包装错误
	wrapped := fmt.Errorf("包装错误: %w", err)
	fmt.Println("fmt.Errorf:", wrapped)

	if result, err := safeDivide(10, 0); err != nil {
		fmt.Println("safeDivide 错误:", err)
	} else {
		fmt.Println("safeDivide 结果:", result)
	}

	riskyOperation()
	fmt.Println("panic 被 recover，程序继续执行")
}
