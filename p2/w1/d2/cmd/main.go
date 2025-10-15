package main


import (
	"w1/d2/config"
)

func main () {
	db, _ := config.ConnectDb()
	defer db.Close()
}