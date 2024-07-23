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

	product2 := Product{
		Name:       "Mouse",
		Price:      100.00,
		CategoryID: category.ID,
	}
	db.Create(&product2)

	db.Create(&SerialNumber{
		Number:    "78910",
		ProductID: product2.ID,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	}
}
