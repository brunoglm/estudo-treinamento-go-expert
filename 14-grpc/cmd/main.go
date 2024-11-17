package main

import (
	"database/sql"
	"graphql/internal/database"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func init() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Cria a tabela inicial
	createCategoriesTableQuery := `
	CREATE TABLE IF NOT EXISTS categories (
		id TEXT PRIMARY KEY, 
		name TEXT NOT NULL, 
		description TEXT NULL
	);
	`
	_, err = db.Exec(createCategoriesTableQuery)
	if err != nil {
		log.Fatalf("Erro ao criar tabela categories: %v\n", err)
	}

	// Cria a tabela inicial
	createCoursesTableQuery := `
	CREATE TABLE IF NOT EXISTS courses (
		id TEXT PRIMARY KEY, 
		name TEXT NOT NULL, 
		description TEXT NULL,
		category_id TEXT,
		FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
	);
	`
	_, err = db.Exec(createCoursesTableQuery)
	if err != nil {
		log.Fatalf("Erro ao criar tabela courses: %v\n", err)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
}
