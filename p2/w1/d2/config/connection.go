package config

import (
	_ "github.com/lib/pq"
	"log"
	"os"
	"fmt"
	"database/sql"
	"github.com/joho/godotenv"
)

func ConnectDb () (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("user")
	password := os.Getenv("password")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")

	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Error connect to database")
	}

	fmt.Println("Connected to database")
	defer db.Close()

	return db, nil
}