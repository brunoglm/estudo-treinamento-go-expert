package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	err = db.Debug().Where("1 = 1").Delete(&Product{}, nil).Error
	if err != nil {
		panic(err)
	}
	err = db.Debug().Where("1 = 1").Delete(&Category{}, nil).Error
	if err != nil {
		panic(err)
	}

	category := Category{
		Name: "Eletronicos",
	}
	db.Create(&category)

	tx := db.Begin()

	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, category.ID).Error
	if err != nil {
		panic(err)
	}

	c.Name = "Eletronicos Super"

	tx.Debug().Save(&c)
	tx.Commit()
}
