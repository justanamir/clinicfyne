package domain

import "time"

type Patient struct {
	ID          string
	FullName    string
	NRIC        string
	DateOfBirth time.Time
	Sex         string

	Phone   string
	Email   string
	Address string

	CreatedAt time.Time
}

// Visit represents a single clinic encounter.
type Visit struct {
	ID        string
	PatientID string

	VisitTime time.Time
	Reason    string

	DoctorID string
	Notes    MedicalNote

	CreatedAt time.Time
}

// MedicalNote represents clinical documentation for a visit.
type MedicalNote struct {
	Subjective string
	Objective  string
	Assessment string
	Plan       string
}

// User represents clinic staff.
type User struct {
	ID       string
	FullName string
	Role     UserRole

	Active bool
}

// UserRole defines staff roles.
type UserRole string

const (
	RoleDoctor       UserRole = "doctor"
	RoleNurse        UserRole = "nurse"
	RoleReceptionist UserRole = "receptionist"
	RoleAdmin        UserRole = "admin"
)
