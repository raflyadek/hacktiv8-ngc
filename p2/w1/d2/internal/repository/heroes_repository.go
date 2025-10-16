package repository

import (
	"database/sql"
	"p2/w1/d2/internal/entity"
)

type HeroesRepository interface {
	GetAllHeroes() ([]entity.Heroes, error)
}

type heroRepo struct {
	db *sql.DB
}

func NewHeroesRepository(db *sql.DB) HeroesRepository {
	return &heroRepo{db}
}

func (h *heroRepo) GetAllHeroes() ([]entity.Heroes, error) {
	rows, err := h.db.Query("SELECT id, name, universe, skill, imageurl FROM heroes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var heroes []entity.Heroes
	for rows.Next() {
		var h entity.Heroes
		if err := rows.Scan(&h.ID, &h.Name, &h.Universe, &h.Skill, &h.ImageURL); err != nil {
			return nil, err
		}
		heroes = append(heroes, h)
	}
	return heroes, nil
}