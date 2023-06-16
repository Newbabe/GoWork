package user

import (
	"GinDemo1/src/conf"
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
}
type List struct {
	Users []User
}

func (u *User) GetId(id int) int {
	return id
}
func Login(password string, name string) bool {
	user := &User{
		Password: password,
		Name:     name,
	}
	user.userLogin()
	fmt.Println("user", user)
	if user.Id > 0 {
		return true
	} else {
		return false
	}
}

func (u *User) userLogin() *User {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	sql := "select id, name, password, email from oa.user where name = ? and password = ?"
	err := conf.DB.QueryRow(sql, u.Name, u.Password).Scan(&u.Id, &u.Name, &u.Password, &u.Email)
	if err != nil {
		return u
	}
	return u
}
