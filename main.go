package main

import (
	"fmt"
	"log"

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

	fmt.Println("Database initialized successfully")
}
