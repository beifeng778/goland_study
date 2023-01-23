package main

//import (
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type Product struct {
//	ID    int
//	Name  string
//	Price float64
//}
//
//func (p *Product) TableName() string {
//	return "product"
//}
//
//func main() {
//	db, err := gorm.Open(mysql.Open("rainyday:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect because")
//	}
//
//	//create
//	db.Create(&Product{ID: 1, Name: "apple", Price: 5.5})
//	var product Product
//	//read
//	db.First(&product, 1)
//	db.First(&product, "id = ?", 1)
//	//update
//	db.Model(&product).Update("Price", 7.5)
//	//updates
//	db.Model(&product).Updates(Product{ID: 2, Price: 7.5})                    //仅支持非零字段
//	db.Model(&product).Updates(map[string]interface{}{"ID": 2, "Price": 7.5}) //支持所有字段
//	//delete
//	db.Delete(&product, 1)
//	db.Clauses()
//}
