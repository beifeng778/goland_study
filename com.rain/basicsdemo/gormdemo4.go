package main

//import (
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type User struct {
//	ID       int    `gorm:"primarykey"`
//	Username string `gorm:"column:username"`
//	Password string `gorm:"column:password"`
//	Age      int    `gorm:"column:age"`
//	Deleted  gorm.DeletedAt
//}
//
//func (u *User) TableName() string {
//	return "user"
//}
//
//func main() {
//	db, err := gorm.Open(mysql.Open("rainyday:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect")
//	}
//	//物理删除
//	//db.Delete(&User{}, 3)
//	//db.Delete(&User{}, "3")
//	//db.Delete(&User{}, []int{1, 2, 3})
//	//db.Where("username LIKE ?", "张%").Delete(&User{})
//	//db.Delete(User{}, "username LIKE ?", "李%")
//	//删除一条
//	u := User{ID: 3}
//	db.Delete(&u)
//	//批量删除
//	db.Where("age = ?", "19").Delete(&User{})
//	//查询时会忽略被软删的记录
//	user := make([]*User, 0)
//	db.Where("age = 19").Find(&user)
//	//查询时间不忽略软删的记录
//	db.Unscoped().Where("age = 19").Find(&user)
//}
