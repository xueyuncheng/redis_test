package model

type User struct {
	Page
	Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) TableName() string {
	return "user"
}
