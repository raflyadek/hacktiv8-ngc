package handler

import (
	"encoding/json"
	
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
	fmt.Println(">>> /villain handler reached")

	villain, err := v.repo.GetAllVillain()
	if err != nil {
		fmt.Print("Error fetching villain", err)
		http.Error(w, "Failed to fetch villain database", http.StatusInternalServerError)
		return
	}
	fmt.Println("Fetched villains:", villain)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villain)
}