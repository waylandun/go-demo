package main

/*
结构体
*/

import "fmt"

// type 关键字有多种用法
// 用法1： 声明一种自定义的数据类型 例：myint  实际是int 的一个bieming
type myint int

// 用法2：定义结构体
type Book struct {
	title  string
	author string
}

func changeBook(book Book) {
	// 尝试修改 --实际不可能起作用 因为是一个值拷贝 不是指针|引用传递
	book.title = "Go 是世界上最好的语言！"
}
func changeBook2(book *Book) {
	// 这个可以成功修改，因为是指针|引用传递 传递的是内存地址
	book.title = "Go Go Go！"
}

func main() {
	book1 := Book{
		title:  "Go Web Programming Language",
		author: "Go Web Programming Language",
	}
	fmt.Println(book1.title)

	var book2 Book
	book2.title = "Python是世界上最好的语言！"
	book2.author = "王五"
	// 字符串格式化输出不能用Println
	//fmt.Println("书名：《%s》 作者：%s", book2.title, book2.author)
	fmt.Printf("书名：《%s》 作者：%s \n", book2.title, book2.author)
	changeBook(book2)
	changeBook2(&book2)
	fmt.Printf("书名：《%s》 作者：%s \n", book2.title, book2.author)
	// %v 是 Go 语言中的格式化动词，用于输出变量的默认格式
	// 对于结构体，%v 会输出字段名和对应的值
	fmt.Printf("%v \n", book2)
}
