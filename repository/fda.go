package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/sergeiten/medilastic"

	log "github.com/sirupsen/logrus"
)

// FDARepository ...
type FDARepository struct {
	DB *sql.DB
}

// NewPermitStatusRepository returns permit status repository
func NewFDARepository(db *sql.DB) Repository {
	return FDARepository{DB: db}
}

// Get return
func (r FDARepository) Get() (map[int]string, error) {
	rows, err := r.DB.Query(`
SELECT id, brand_name, company_name, device_description, gmdn_pt_name, gmdn_pt_definition, product_code, product_code_name
FROM fda_products`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := map[int]string{}

	for rows.Next() {
		item := medilastic.Fda{}

		err := rows.Scan(&item.ID, &item.BrandName, &item.CompanyName, &item.DeviceDescription, &item.GmdnPtName, &item.GmdnPtDefinition, &item.ProductCode, &item.ProductCodeName)
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
