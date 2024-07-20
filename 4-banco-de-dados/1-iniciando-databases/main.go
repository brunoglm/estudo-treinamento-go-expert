package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("DB conectado com sucesso!")

	p := NewProduct("name1", 10.11)
	err = insertProduct(db, p)
	if err != nil {
		panic(err)
	}

	p.Name = "Name mudado"
	p.Price = 20.22
	err = updateProduct(db, p)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("insert into products (id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(p.Id, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(p.Name, p.Price, p.Id)
	if err != nil {
		return err
	}
	return nil
}
