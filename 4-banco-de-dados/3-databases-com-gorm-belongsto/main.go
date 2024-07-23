package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Name string
	gorm.Model
}

type Product struct {
	Name       string
	Price      float64
	CategoryID uint
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	category := Category{
		Name: "Eletronicos",
	}
	db.Create(&category)

	product := Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: category.ID,
	}
	db.Create(&product)

	product2 := Product{
		Name:       "Mouse",
		Price:      100.00,
		CategoryID: category.ID,
	}
	db.Create(&product2)

	var products []Product
	db.Preload("Category").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name)
	}
}
