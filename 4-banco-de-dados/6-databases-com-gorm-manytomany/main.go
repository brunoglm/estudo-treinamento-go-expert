package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Product struct {
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
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

	category2 := Category{
		Name: "Cozinha",
	}
	db.Create(&category2)

	product := Product{
		Name:       "Notebook",
		Price:      1000.00,
		Categories: []Category{category, category2},
	}
	db.Create(&product)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println("- ", p.Name)
		}
	}
}
