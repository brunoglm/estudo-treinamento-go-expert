package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	insertProduct(db)
	insertManyProducts(db)
	selectOneProduct(db)
	selectAllProducts(db)
	alterandoDados(db)
	deletandoDados(db)
}

func insertProduct(db *gorm.DB) {
	db.Create(&Product{
		Name:  "name 1",
		Price: 11.22,
	})
}

func insertManyProducts(db *gorm.DB) {
	db.Create(&[]Product{
		{Name: "name 2", Price: 55.33},
		{Name: "name 3", Price: 55.33},
	})
}

func selectOneProduct(db *gorm.DB) {
	var product Product
	db.First(&product, 1)
	fmt.Println("product: ", product)

	var product2 Product
	db.First(&product2, "name = ?", "name 3")
	fmt.Println("product2: ", product2)
}

func selectAllProducts(db *gorm.DB) {
	var products []Product
	db.Find(&products)
	fmt.Println("products all: ", products)

	var products2 []Product
	db.Limit(2).Find(&products2)
	fmt.Println("products2 limit: ", products2)

	var products3 []Product
	db.Limit(2).Offset(1).Find(&products3)
	fmt.Println("products3 limit e offset: ", products3)

	var products4 []Product
	db.Where("name = ?", "name 3").Find(&products4)
	fmt.Println("products4 where: ", products4)

	var products5 []Product
	db.Where("name LIKE ?", "%name 3%").Find(&products5)
	fmt.Println("products5 where e like: ", products5)
}

func alterandoDados(db *gorm.DB) {
	var product Product
	db.First(&product, 1)
	fmt.Println("product à ser alterado: ", product)

	product.Name = "produto alterado"

	db.Save(&product)

	var product2 Product
	db.First(&product2, 1)
	fmt.Println("product alterado: ", product2)
}

func deletandoDados(db *gorm.DB) {
	var product Product
	db.First(&product, 1)
	fmt.Println("product à ser deletado: ", product)

	db.Delete(&product)
}
