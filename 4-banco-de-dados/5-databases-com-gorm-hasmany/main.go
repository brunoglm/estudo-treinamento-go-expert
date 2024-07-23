package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Name     string
	Products []Product
	gorm.Model
}

type Product struct {
	Name         string
	Price        float64
	CategoryID   uint
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	Number    string
	ProductID uint
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

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

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println("- ", p.Name, ", Serial Number: ", p.SerialNumber.Number)
		}
	}
}
