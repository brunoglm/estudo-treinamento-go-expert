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

	product, err := selectProduct(db, p.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product: ", product)

	products, err := selectAllProduct(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Products: ", products)

	err = deleteProduct(db, p.Id)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("insert into products (id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
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
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Price, p.Id)
	if err != nil {
		return err
	}
	return nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	p := Product{}
	err = stmt.QueryRow(id).Scan(&p.Id, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProduct(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
