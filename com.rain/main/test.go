package main

import (
	"fmt"
)

type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user{name: "zheng", password: "123"}
	a.resetPassword("456")
	fmt.Println(a.checkPassword("456")) //true
}
