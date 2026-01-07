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

	if err := patientRepo.Create(patient); err != nil {
		log.Fatal(err)
	}

	found, err := patientRepo.GetByID("p001")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Patient loaded: %+v\n", found)
}
