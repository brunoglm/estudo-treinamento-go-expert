package main

import (
	"database/sql"
	"grpc/internal/database"
	"grpc/internal/pb"
	"grpc/internal/service"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	categoryService := service.NewCategoryService(categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
