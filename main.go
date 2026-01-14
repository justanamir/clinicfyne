package main

import (
	"fmt"
	"log"
	"time"

	"clinicfyne/domain"
	"clinicfyne/storage"
)

func main() {
	db, err := storage.OpenDB("clinic.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.RunMigrations(db); err != nil {
		log.Fatal(err)
	}

	patientRepo := storage.NewPatientRepository(db)
	visitRepo := storage.NewVisitRepository(db)

	patient := domain.Patient{
		ID:          "p001",
		FullName:    "Ahmad Bin Ali",
		NRIC:        "900101-14-5678",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Sex:         "M",
		Phone:       "012-3456789",
		Email:       "ahmad@example.com",
		Address:     "Kuala Lumpur",
		CreatedAt:   time.Now(),
	}

	_ = patientRepo.Create(patient)

	visit := domain.Visit{
		ID:        "v001",
		PatientID: patient.ID,
		VisitTime: time.Now(),
		Reason:    "Fever and cough",
		DoctorID:  "d001",
		Notes: domain.MedicalNote{
			Subjective: "Patient complains of fever for 3 days",
			Objective:  "Temperature 38.5",
			Assessment: "Upper respiratory tract infection",
			Plan:       "Paracetamol, rest, fluids",
		},
		CreatedAt: time.Now(),
	}

	if err := visitRepo.Create(visit); err != nil {
		log.Fatal(err)
	}

	history, err := visitRepo.ListByPatientID(patient.ID)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range history {
		fmt.Printf("Visit on %s: %s\n", v.VisitTime.Format("2006-01-02"), v.Reason)
	}

}
