package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/sergeiten/medilastic"

	log "github.com/sirupsen/logrus"
)

// PermitStatusRepository ...
type PermitStatusRepository struct {
	DB *sql.DB
}

// NewPermitStatusRepository returns permit status repository
func NewPermitStatusRepository(db *sql.DB) Repository {
	return PermitStatusRepository{DB: db}
}

// Get return
func (r PermitStatusRepository) Get() (map[int]string, error) {
	rows, err := r.DB.Query(`
SELECT id, prduct, entrps, prduct_prmisn_no, mea_class_no, type_name, use_purps
FROM permit_status`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := map[int]string{}

	for rows.Next() {
		item := medilastic.PermitStatus{}

		err := rows.Scan(&item.ID, &item.Prduct, &item.Entrps, &item.PrductPrmisnNo, &item.MeaClassNo, &item.TypeName, &item.UsePurps)
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
