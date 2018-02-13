package repository

import "database/sql"

// Repository interface
type Repository interface {
	Get() (map[int]string, error)
}

// NewRepository returns specific repository by passed name
func NewRepository(name string, db *sql.DB) Repository {
	switch name {
	case "permit_status":
		return NewPermitStatusRepository(db)
	case "fda":
		return NewFDARepository(db)
	}

	return nil
}
