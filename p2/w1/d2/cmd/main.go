package main

import (
	"fmt"
	"log"
	"net/http"
	"p2/w1/d2/config"
	"p2/w1/d2/internal/handler"
	"p2/w1/d2/internal/repository"
)

func main () {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatal("Error connect to database")
	}
	defer db.Close()

	
	//repo
	heroRepo := repository.NewHeroesRepository(db)
	villainRepo := repository.NewVillainRepository(db)
	
	//handler
	heroHandler := handler.NewHeroHandler(heroRepo)
	villainHandler := handler.NewVillainHandler(villainRepo)
	
	//route
	http.HandleFunc("/heroes", heroHandler.GetAllHeroes)
	http.HandleFunc("/villain", villainHandler.GetAllVillain)

	// er := config.ConnectServer()
	// if er != nil {
	// 	log.Print("Error conect to server", er)
	// }
	log.Print("Server started at port :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err.Error())
	}

}