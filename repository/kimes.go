package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/sergeiten/medilastic"

	log "github.com/sirupsen/logrus"
)

// KimesRepository ...
type KimesRepository struct {
	DB *sql.DB
}

// NewKimesRepository returns kimes repository
func NewKimesRepository(db *sql.DB) Repository {
	return KimesRepository{DB: db}
}

// Get return
func (r KimesRepository) Get() (map[int]string, error) {
	rows, err := r.DB.Query(`
SELECT id, name, model, country, manufacture, specification, description, category, subcategory
FROM kimes_products ORDER BY id`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := map[int]string{}

	for rows.Next() {
		item := medilastic.Kimes{}

		err := rows.Scan(&item.ID, &item.Name, &item.Model, &item.Country, &item.Manufacture, &item.Specification, &item.Description, &item.Category, &item.Subcategory)
		if err != nil {
			log.WithError(err).Error("failed to scan rows")
		}

		dat, err := json.Marshal(item)
		if err != nil {
			log.WithError(err).Error("failed to marshal json")
		}

		items[item.ID] = string(dat)
	}

	return items, nil
}
