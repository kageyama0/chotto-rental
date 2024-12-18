package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Case{},
		&Application{},
		&Matching{},
		&Review{},
		&Session{},
	)
}

// User ユーザーモデル
// @Description ユーザー情報
type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	DisplayName  string    `gorm:"not null"`
	TrustScore   float64   `gorm:"default:1.0"`
	NoShowCount  int       `gorm:"default:0"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// Case 案件モデル
// @Description 案件情報
type Case struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID          uuid.UUID `gorm:"not null"` // 案件の依頼者
	Title           string    `gorm:"not null"`
	Description     string    `gorm:"not null"`
	Reward          int       `gorm:"not null"`
	Location        string    `gorm:"not null"`
	ScheduledDate   time.Time `gorm:"not null"`
	DurationMinutes int       `gorm:"not null"`
	Status          string    `gorm:"default:open"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User            User      `gorm:"foreignkey:UserID"`
}

// Application 応募モデル
// @description 応募情報
type Application struct {
	gorm.Model
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

// @description マッチング情報
type Matching struct {
	gorm.Model
	ID                          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CaseID                      uuid.UUID `gorm:"not null"`
	RequesterID                 uuid.UUID `gorm:"not null"` // 案件の依頼者
	HelperID                    uuid.UUID `gorm:"not null"` // 案件を受けたユーザー
	MeetingLocation             string    `gorm:"not null"`
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

// Review レビューモデル
// @description レビュー情報
type Review struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	MatchingID     uuid.UUID `gorm:"not null"`
	ReviewerID     uuid.UUID `gorm:"not null"` // レビューを書いたユーザー
	ReviewedUserID uuid.UUID `gorm:"not null"` // レビューを書かれたユーザー
	Score          int       `gorm:"not null;check:score >= 1 AND score <= 5"`
	Comment        string
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Matching       Matching  `gorm:"foreignkey:MatchingID"`
	Reviewer       User      `gorm:"foreignkey:ReviewerID"`
	ReviewedUser   User      `gorm:"foreignkey:ReviewedUserID"`
}

type Session struct {
	gorm.Model
	ID             uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID         uuid.UUID  `gorm:"type:varchar(36);not null"`
	DeviceInfo     DeviceInfo `gorm:"type:jsonb;not null"`
	CreatedAt      time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	LastAccessedAt time.Time  `gorm:"not null"`
	ExpiresAt      time.Time  `gorm:"not null"`
	IsValid        bool       `gorm:"not null;default:true"`

	User User `gorm:"foreignKey:UserID"`
}

type DeviceInfo struct {
	UserAgent    string `json:"user_agent"`
	IP           string `json:"ip"`
	ClientName   string `json:"client_name"`             // モバイルアプリ名など
	DeviceID     string `json:"device_id,omitempty"`     // デバイス識別子（オプション）
	LastLocation string `json:"last_location,omitempty"` // 最後のアクセス位置（オプション）
}
