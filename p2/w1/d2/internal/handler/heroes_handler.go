package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"p2/w1/d2/internal/repository"
	"fmt"
)

type HeroHandler struct {
	repo repository.HeroesRepository
}

func NewHeroHandler(repo repository.HeroesRepository) *HeroHandler {
	return &HeroHandler{repo}
}

func (h *HeroHandler) GetAllHeroes(w http.ResponseWriter, r *http.Request) {
	heroes, err := h.repo.GetAllHeroes()
	if err != nil {
		log.Print("Error fetching heroes", err)
		http.Error(w, "Failed to fetch heroes database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Fetched villains:", heroes)
	json.NewEncoder(w).Encode(heroes)
}