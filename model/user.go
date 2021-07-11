package model

import "gorm.io/gorm"

type AddUserReq struct {
	Name string
	Age  int
}

type ListUserReq struct {
	Page
	Name string
}

type ListUserResp struct {
	ID   int
	Name string
	Age  int
}

type User struct {
	gorm.Model
	Name string
	Age  int
}

func (u *User) TableName() string {
	return "user"
}
