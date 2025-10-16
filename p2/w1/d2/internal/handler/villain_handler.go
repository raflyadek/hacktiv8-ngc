package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"p2/w1/d2/internal/repository"
	"fmt"
)

type VillainHandler struct {
	repo repository.VillainRepository
}

func NewVillainHandler(repo repository.VillainRepository) *VillainHandler {
	return &VillainHandler{repo}
}

func (v *VillainHandler) GetAllVillain(w http.ResponseWriter, r *http.Request) {
	villain, err := v.repo.GetAllVillain()
	if err != nil {
		log.Print("Error fetching villain", err)
		http.Error(w, "Failed to fetch villain database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Fetched villains:", villain)
	json.NewEncoder(w).Encode(villain)
}