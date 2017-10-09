package models

type Auth struct {
	Id      int
	IsLogin bool
	Type    string
	User
}
