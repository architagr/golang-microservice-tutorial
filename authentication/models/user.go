package models

type User struct {
	Id       int
	Name     string
	UserName string
	Password string
	Roles    []int
}
