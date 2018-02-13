package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/sergeiten/medilastic"

	log "github.com/sirupsen/logrus"
)

// PmaRepository ...
type PmaRepository struct {
	DB *sql.DB
}

// NewPmaRepository returns pma repository
func NewPmaRepository(db *sql.DB) Repository {
	return PmaRepository{DB: db}
}

// Get return
func (r PmaRepository) Get() (map[int]string, error) {
	rows, err := r.DB.Query(`
SELECT id, applicant, genericname, tradename
FROM pma_products`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := map[int]string{}

	for rows.Next() {
		item := medilastic.Pma{}

		err := rows.Scan(&item.ID, &item.Applicant, &item.GenericName, &item.TradeName)
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
