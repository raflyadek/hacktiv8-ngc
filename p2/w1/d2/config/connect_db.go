package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"database/sql"
)

func ConnectDb() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	user := os.Getenv("user")
	password := os.Getenv("password")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, dbname)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("error Connect to database", err)
	}

	fmt.Println("Success connect to database!")
	return db, nil 
}