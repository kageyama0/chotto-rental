package application_repository

import (
	"github.com/google/uuid"
)

// あるユーザーが、ある案件に応募しているかどうかを調べる
func (r *ApplicationRepository) FindByCaseIDAndApplicantID(caseID uuid.UUID, applicantID uuid.UUID) (Application, error) {
	var application Application

	if err := r.db.First(&application, "case_id = ? AND applicant_id = ?", caseID, applicantID).Error; err != nil {
		return application, err
	}

	return application, nil
}

// 応募のIDを使用して、案件を紐づけた応募を取得する
func (r *ApplicationRepository) FindByIDWithCase(id uuid.UUID) (Application, error) {
	var application Application

	err := r.db.Preload("Case").First(&application, "id = ?", id).Error

	if err != nil {
		return application, err
	}

	return application, nil
}

// すべての応募とそれに紐づく案件を取得する
func (r *ApplicationRepository) FindAllByIDWithCase(userID *uuid.UUID) ([]Application, error) {
	var applications []Application

	err := r.db.Preload("Case").Find(&applications, "applicant_id = ?", userID).Error

	if err != nil {
		return applications, err
	}

	return applications, nil
}
