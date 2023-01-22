package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	id       int
	username string
	password string
}

func (user *User) TableName() string {
	return "user"
}

func main() {
	dsn := "rainyday:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	user := User{username: "王五", password: "123456"}

	result := db.Model(&user) // pass pointer of data to Create

	fmt.Println(user.id)
	fmt.Println(user.username)       // returns inserted data's primary key
	fmt.Println(result.Error)        // returns error
	fmt.Println(result.RowsAffected) // returns inserted records count
}
