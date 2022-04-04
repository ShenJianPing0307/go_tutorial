package model

// 定义一个私有结体
type user struct {
	name string
	Age  int
}

// 工厂方法，它包可以调用该方法
func NewUser(name string, age int) *user {

	return &user{
		name: name,
		Age:  age,
	}

}

// 私有字段name工厂方法
func GetName(u *user) string {
	return u.name
}
