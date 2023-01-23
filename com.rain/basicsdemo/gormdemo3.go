package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"primarykey"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Age      int    `gorm:"column:age"`
}

func (p *User) TableName() string {
	return "user"
}

func main() {
	db, err := gorm.Open(mysql.Open("rainyday:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect")
	}
	//查询第一条数据，按照主键升序，查询不到返回ErrRecordNotFound
	user := &User{}
	db.First(user)
	//查询多条数据
	users := make([]*User, 0)
	res := db.Where("age > 20").Find(&users)
	fmt.Println(res.RowsAffected)
	fmt.Println(res.Error)
	//IN操作
	db.Where("username IN ?", []string{"张三", "张四"}).Find(&users)
	//LIKE操作
	db.Where("username LIKE ?", "张%").Find(&user)
	//AND操作
	db.Where("username = ? AND age >= ?", "张三", "19").Find(&users)

	//结构体查询方式
	db.Where(&User{Username: "张三", Age: 20}).Find(&users)
	//map查询方式
	db.Where(map[string]interface{}{"Username": "张三", "Age": 20}).Find(&users)
}
