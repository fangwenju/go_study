package common

import (
	"fmt"
)

const (
	user_name = "fangwenju"
	password  = "123456"
)

type User struct {
	user_name string
	password  string
	status    bool
}

func (u *User) login() {
	fmt.Print("请输入登录名：")
	fmt.Scan(&u.user_name)
	fmt.Print("请输入密码：")
	fmt.Scan(&u.password)

	if u.user_name == user_name && password == u.password {
		fmt.Println("\n>>>>>>>>>>>>> 验证成功 <<<<<<<<<<<<<<")
		fmt.Println(">>>>>> 欢迎使用练习查询系统!!!! <<<<<<<<")
	} else {
		fmt.Println(">>>>>>>>>>>>> 验证失败 <<<<<<<<<<<<<<<<")
		u.login()
	}
}

func (u *User) Verify() {
	if u.status != true {
		u.login()
	}

	u.status = true
}

func (u *User) Logout() {
	u.status = false
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>> END <<<<<<<<<<<<<<<<<<<<<<<<<<<")
}
