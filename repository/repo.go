package repository

import (
	"criminal_report/entity"
	"database/sql"
	"fmt"
)

type CriminalRepository struct {
	db *sql.DB
}

func NewCriminalRepository(db *sql.DB) *CriminalRepository {
	return &CriminalRepository{db}
}

func (repo *CriminalRepository) GetAllCriminals() ([]entity.CriminalReport, error) {
	var criminals []entity.CriminalReport
	query := "SELECT id, heroid, villainid, description, occurrence FROM criminal_reports"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var criminal entity.CriminalReport
		err := rows.Scan(&criminal.ID, &criminal.HeroID, &criminal.VillainID, &criminal.Description, &criminal.Occurrence)
		if err != nil {
			return nil, err
		}
		criminals = append(criminals, criminal)
	}

	return criminals, nil
}

func (repo *CriminalRepository) GetCriminalByID(id int) (entity.CriminalReport, error) {
	var criminal entity.CriminalReport
	query := "SELECT id, heroid, villainid, description, occurrence FROM criminal_reports WHERE id = ?"
	err := repo.db.QueryRow(query, id).Scan(&criminal.ID, &criminal.HeroID, &criminal.VillainID, &criminal.Description, &criminal.Occurrence)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.CriminalReport{}, fmt.Errorf("criminal with ID %d not found", id)
		}
		return entity.CriminalReport{}, err
	}
	return criminal, nil
}

func (repo *CriminalRepository) CreateCriminal(criminal entity.CriminalReport) (int, error) {
	query := "INSERT INTO criminal_reports (heroid, villainid, description, occurrence) VALUES (?, ?, ?, ?)"
	result, err := repo.db.Exec(query, criminal.HeroID, criminal.VillainID, criminal.Description, criminal.Occurrence)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}

func (repo *CriminalRepository) UpdateCriminal(id int, criminal entity.CriminalReport) error {
	query := "UPDATE criminal_reports SET heroid = ?, villainid = ?, description = ?, occurrence = ? WHERE id = ?"
	_, err := repo.db.Exec(query, criminal.HeroID, criminal.VillainID, criminal.Description, criminal.Occurrence, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CriminalRepository) DeleteCriminal(id int) error {
	query := "DELETE FROM criminal_reports WHERE id = ?"
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
