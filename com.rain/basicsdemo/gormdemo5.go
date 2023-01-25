package main

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `gorm:"primarykey"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Age      int    `gorm:"column:age"`
}

type Email struct {
	ID       int
	Username string
	Email    string
}

func (p *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age < 0 {
		return errors.New("can't save invalid data")
	}
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Create(&Email{ID: u.ID, Email: u.Username + "@qq.com"}).Error
}
func main() {
	db, err := gorm.Open(mysql.Open("rainyday:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic("failed to connect")
	}
	//开始事务
	tx := db.Begin()
	/*
		在事务中执行的db操作，这里使用tx而不是db
	*/
	//回滚事务
	if err = tx.Create(&User{Username: "张三"}).Error; err != nil {
		tx.Rollback()
		return
	}
	//提交事务
	tx.Commit()

	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&User{Username: "张三"}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return
	}
}
