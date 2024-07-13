package pojo

// 在结构体中，无论是结构体名称还是字段名称，大写都表示公有，可以被调用 . 小写都表示私有，只能自己使用

type User struct {
	Name string
	Age  int
	Sex  string
}

type Cat struct {
	Color  string
	Weight int
}
