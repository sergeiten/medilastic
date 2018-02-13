package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/sergeiten/medilastic"

	log "github.com/sirupsen/logrus"
)

// PasRepository ...
type PasRepository struct {
	DB *sql.DB
}

// NewPasRepository returns pas repository
func NewPasRepository(db *sql.DB) Repository {
	return PasRepository{DB: db}
}

// Get return
func (r PasRepository) Get() (map[int]string, error) {
	rows, err := r.DB.Query(`
SELECT id, application_name, device_name, medical_speciality, study_name, study_design_description
FROM pas_products`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := map[int]string{}

	for rows.Next() {
		item := medilastic.Pas{}

		err := rows.Scan(&item.ID, &item.ApplicationName, &item.DeviceName, &item.MedicalSpeciality, &item.StudyName, &item.StudyDesignDescription)
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
