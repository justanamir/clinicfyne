package storage

import (
	"clinicfyne/domain"
	"database/sql"
	"time"
)

type PatientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

// Create insert a new patient record.
func (r *PatientRepository) Create(p domain.Patient) error {
	query := `
	INSERT INTO patients (
	id,
	full_name,
	nric,
	date_of_birth,
	sex,
	phone,
	email,
	address,
	created_at
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := r.db.Exec(
		query,
		p.ID,
		p.FullName,
		p.NRIC,
		p.DateOfBirth.Format(time.RFC3339),
		p.Sex,
		p.Phone,
		p.Email,
		p.Address,
		p.CreatedAt.Format(time.RFC3339),
	)

	return err
}

func (r *PatientRepository) GetByID(id string) (*domain.Patient, error) {
	query := `
	SELECT
	id,
	full_name,
	nric,
	date_of_birth,
	sex,
	phone,
	email,
	address,
	created_at
	FROM patients
	WHERE id = ?;
	`

	row := r.db.QueryRow(query, id)

	var p domain.Patient
	var dob, created string

	err := row.Scan(
		&p.ID,
		&p.FullName,
		&p.NRIC,
		&dob,
		&p.Sex,
		&p.Phone,
		&p.Email,
		&p.Address,
		&created,
	)

	if err != nil {
		return nil, err
	}

	p.DateOfBirth, _ = time.Parse(time.RFC3339, dob)
	p.CreatedAt, _ = time.Parse(time.RFC3339, created)

	return &p, nil
}
