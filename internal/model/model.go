package model

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	DisplayName  string    `gorm:"not null"`
	TrustScore   float64   `gorm:"default:1.0"`
	NoShowCount  int       `gorm:"default:0"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Case struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID         uuid.UUID `gorm:"not null"`
	Title          string    `gorm:"not null"`
	Description    string    `gorm:"not null"`
	Reward         int       `gorm:"not null"`
	Location       string    `gorm:"not null"`
	ScheduledDate  time.Time `gorm:"not null"`
	DurationMinutes int      `gorm:"not null"`
	Status         string    `gorm:"default:open"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User           User      `gorm:"foreignkey:UserID"`
}

type Application struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CaseID      uuid.UUID `gorm:"not null"`
	ApplicantID uuid.UUID `gorm:"not null"`
	Status      string    `gorm:"default:pending"`
	Message     string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Case        Case      `gorm:"foreignkey:CaseID"`
	Applicant   User      `gorm:"foreignkey:ApplicantID"`
}

type Matching struct {
	ID                          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CaseID                      uuid.UUID `gorm:"not null"`
	RequesterID                 uuid.UUID `gorm:"not null"`
	HelperID                    uuid.UUID `gorm:"not null"`
	MeetingLocation            string    `gorm:"not null"`
	ArrivalConfirmedByRequester bool      `gorm:"default:false"`
	ArrivalConfirmedByHelper    bool      `gorm:"default:false"`
	ArrivalConfirmationDeadline time.Time `gorm:"not null"`
	Status                      string    `gorm:"default:active"`
	CreatedAt                   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Case                        Case      `gorm:"foreignkey:CaseID"`
	Requester                   User      `gorm:"foreignkey:RequesterID"`
	Helper                      User      `gorm:"foreignkey:HelperID"`
}

type Review struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	MatchingID     uuid.UUID `gorm:"not null"`
	ReviewerID     uuid.UUID `gorm:"not null"`
	ReviewedUserID uuid.UUID `gorm:"not null"`
	Score          int       `gorm:"not null;check:score >= 1 AND score <= 5"`
	Comment        string
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Matching       Matching  `gorm:"foreignkey:MatchingID"`
	Reviewer       User      `gorm:"foreignkey:ReviewerID"`
	ReviewedUser   User      `gorm:"foreignkey:ReviewedUserID"`
}
