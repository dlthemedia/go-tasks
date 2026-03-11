package main

type InfoContactInfo struct {
	Email string
	Phone string
}

type Student struct {
	ID   string
	Name string
	InfoContactInfo
}

type Course struct {
	Code        string
	Title       string
	CreditHours int
}

type EnrollmentStatus string

const (
	StatusActive    EnrollmentStatus = "active"
	StatusCompleted EnrollmentStatus = "completed"
	StatusDropped   EnrollmentStatus = "dropped"
)

type CourseEnrollment struct {
	EnrollmentID string
	Status       EnrollmentStatus
	Grade        int
	Student      Student
	Course       Course
}
