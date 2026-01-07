package storage

import "database/sql"

func RunMigrations(db *sql.DB) error {
	stmts := []string{
		`
		CREATE TABLE IF NOT EXISTS patients (
			id TEXT PRIMARY KEY,
			full_name TEXT NOT NULL,
			nric TEXT NOT NULL,
			date_of_birth TEXT NOT NULL,
			sex TEXT NOT NULL,
			phone TEXT,
			email TEXT,
			address TEXT,
			created_at TEXT NOT NULL
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			full_name TEXT NOT NULL,
			role TEXT NOT NULL,
			active INTEGER NOT NULL
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS visits (
			id TEXT PRIMARY KEY,
			patient_id TEXT NOT NULL,
			visit_time TEXT NOT NULL,
			reason TEXT,
			doctor_id TEXT,
			subjective TEXT,
			objective TEXT,
			assessment TEXT,
			plan TEXT,
			created_at TEXT NOT NULL,
			FOREIGN KEY(patient_id) REFERENCES patients(id)
		);
		`,
	}

	for _, stmt := range stmts {
		if _, err := db.Exec(stmt); err != nil {
			return err
		}
	}

	return nil
}
