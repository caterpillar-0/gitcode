package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate(迁移) the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	//read
	var product Product
	db.First(&product, 1)                 //根据整形主键查找
	db.First(&product, "code = ?", "D42") //查找code字段为D42的记录

	//update,将product的price更新为200
	db.Model(&product).Update("Pice", 200)
	//update更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) //仅更新非零字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//delete 删除product
	db.Delete(&product, 1)

}
