// 教学内容: 条件语句 if/else、switch 类型、for 循环、range 遍历、defer 延迟执行与 goto。
package main

import "fmt"

func main() {
	// if 支持在条件前初始化语句
	if score := 85; score >= 90 {
		fmt.Println("成绩优秀")
	} else if score >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}

	// switch 不需要 break，支持类型开关和 fallthrough
	day := 3
	switch day {
	case 1, 2, 3:
		fmt.Println("工作日的前半段")
	case 4, 5:
		fmt.Println("工作日的后半段")
	default:
		fmt.Println("周末")
	}

	// 类型 switch
	describe := func(v interface{}) {
		switch value := v.(type) {
		case int:
			fmt.Println("整型:", value)
		case string:
			fmt.Println("字符串:", value)
		default:
			fmt.Println("未知类型")
		}
	}
	describe("hello")

	// for 的三种形式
	for i := 0; i < 3; i++ {
		fmt.Printf("经典 for i=%d\n", i)
	}

	n := 3
	for n > 0 {
		fmt.Printf("条件 for n=%d\n", n)
		n--
	}

	// 无限循环配合 break
	counter := 0
	for {
		counter++
		if counter == 2 {
			fmt.Println("continue 跳过")
			continue
		}
		fmt.Printf("无限循环: counter=%d\n", counter)
		if counter >= 3 {
			break
		}
	}

	// range 遍历集合
	langs := []string{"Go", "Rust", "Python"}
	for index, value := range langs {
		fmt.Printf("range: index=%d value=%s\n", index, value)
	}

	// defer 延迟执行: 先进后出
	defer fmt.Println("defer 最后执行")
	fmt.Println("主逻辑执行中")

	// goto 跳转
	goto Label

Label:
	fmt.Println("通过 goto 跳转到这里")
}
