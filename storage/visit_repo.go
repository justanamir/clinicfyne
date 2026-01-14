package storage

import (
	"clinicfyne/domain"
	"database/sql"
	"time"
)

type VisitRepository struct {
	db *sql.DB
}

func NewVisitRepository(db *sql.DB) *VisitRepository {
	return &VisitRepository{db: db}
}

// Create inserts a new visit record.
func (r *VisitRepository) Create(v domain.Visit) error {
	query := `
	INSERT INTO visits (
	id,
	patient_id,
	visit_time,
	reason,
	doctor_id,
	subjective,
	objective,
	assessment,
	plan,
	created_at 
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := r.db.Exec(
		query,
		v.ID,
		v.PatientID,
		v.VisitTime.Format(time.RFC3339),
		v.Reason,
		v.DoctorID,
		v.Notes.Subjective,
		v.Notes.Objective,
		v.Notes.Assessment,
		v.Notes.Plan,
		v.CreatedAt.Format(time.RFC3339),
	)

	return err
}

// ListByPatientID returns all visits for a patient, newest first.
func (r *VisitRepository) ListByPatientID(patientID string) ([]domain.Visit, error) {
	query := `
	SELECT 
	id,
	patient_id,
	visit_time,
	reason,
	doctor_id,
	subjective,
	objective,
	assessment,
	plan,
	created_at
	FROM visits
	WHERE patient_id = ?
	ORDER BY visit_time DESC;
	`

	rows, err := r.db.Query(query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visits []domain.Visit

	for rows.Next() {
		var v domain.Visit
		var visitTime, created string

		err := rows.Scan(
			&v.ID,
			&v.PatientID,
			&visitTime,
			&v.Reason,
			&v.DoctorID,
			&v.Notes.Subjective,
			&v.Notes.Objective,
			&v.Notes.Assessment,
			&v.Notes.Plan,
			&created,
		)
		if err != nil {
			return nil, err
		}

		v.VisitTime, _ = time.Parse(time.RFC3339, visitTime)
		v.CreatedAt, _ = time.Parse(time.RFC3339, created)

		visits = append(visits, v)
	}

	return visits, nil
}
