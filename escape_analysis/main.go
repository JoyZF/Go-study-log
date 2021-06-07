package main

type User struct {
	ID     int64
	Name   string
	Avatar string
}

func GetUserInfo() *User {
	return &User{ID: 123456,Name: "张三",Avatar: "https://baidu.com"}
}

func main() {
	_ = GetUserInfo()
}
