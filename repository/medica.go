package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/sergeiten/medilastic"

	log "github.com/sirupsen/logrus"
)

// MedicaRepository ...
type MedicaRepository struct {
	DB *sql.DB
}

// NewMedicaRepository returns medica repository
func NewMedicaRepository(db *sql.DB) Repository {
	return MedicaRepository{DB: db}
}

// Get return
func (r MedicaRepository) Get() (map[int]string, error) {
	rows, err := r.DB.Query(`
SELECT id, title, description, company_title, company_description
FROM medica_products`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := map[int]string{}

	for rows.Next() {
		item := medilastic.Medica{}

		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.CompanyTitle, &item.CompanyDescription)
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
