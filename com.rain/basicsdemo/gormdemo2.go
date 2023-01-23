package main

//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type Product struct {
//	ID    int     `gorm:"primarykey"`
//	Name  string  `gorm:"column:name"`
//	Price float64 `gorm:"column:price"`
//}
//
//func (p *Product) TableName() string {
//	return "product"
//}
//
//func main() {
//	db, err := gorm.Open(mysql.Open("rainyday:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect")
//	}
//	//创建一条数据
//	p := &Product{Name: "pear"}
//	res := db.Create(p)
//	fmt.Println(res.Error)
//	fmt.Println(p.ID)
//	//创建多条数据
//	products := []*Product{{Name: "orange"}, {Name: "banana"}}
//	res = db.Create(products)
//	fmt.Println(res.Error)
//	for _, p := range products {
//		fmt.Println(p.ID)
//	}
//}
