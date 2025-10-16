package repository

import (
	"database/sql"
	"p2/w1/d2/internal/entity"
)

type VillainRepository interface {
	GetAllVillain() ([]entity.Villain, error)
}

type villainRepo struct {
	db *sql.DB
}

func NewVillainRepository(db *sql.DB) VillainRepository {
	return &villainRepo{db}
}

func (v *villainRepo) GetAllVillain() ([]entity.Villain, error) {
	rows, err := v.db.Query(`SELECT id, name, universe, Imageurl FROM villain`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var villain []entity.Villain
	for rows.Next() {
		var v entity.Villain
		if err := rows.Scan(&v.ID, &v.Name, &v.Universe, &v.ImageURL); err != nil {
			return nil, err
		}
		villain = append(villain, v)
	}
	return villain, nil
}