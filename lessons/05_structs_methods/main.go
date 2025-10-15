// 教学内容: 结构体定义、字段标签、构造函数、方法集、值接收者与指针接收者、匿名字段与嵌套。
package main

import "fmt"

// User 描述一个用户，包含导出与未导出字段
type User struct {
	ID    int
	Name  string
	email string // 未导出字段，仅包内可访问
}

// NewUser 作为构造函数使用
func NewUser(id int, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		email: email,
	}
}

// Greeting 值接收者方法，不会修改原始数据
func (u User) Greeting() string {
	return fmt.Sprintf("你好，我是 %s", u.Name)
}

// UpdateEmail 指针接收者方法，可修改原始值
func (u *User) UpdateEmail(newEmail string) {
	u.email = newEmail
}

// Admin 通过匿名字段嵌套 User
type Admin struct {
	User
	Roles []string
}

func main() {
	user := NewUser(1, "Alice", "alice@example.com")
	fmt.Println(user.Greeting())

	user.UpdateEmail("alice@new.com")
	fmt.Printf("更新后的邮箱: %s\n", user.email)

	admin := Admin{
		User:  *user,
		Roles: []string{"manage_users", "view_reports"},
	}

	// 可直接访问嵌入字段的方法与属性
	fmt.Println("管理员问候语:", admin.Greeting())
	fmt.Println("管理员角色:", admin.Roles)
}
