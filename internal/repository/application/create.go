package application_repository

import (
	"github.com/kageyama0/chotto-rental/internal/model"
)

// 案件に対する応募を作成する
func (r *ApplicationRepository) Create(application *model.Application) error {
	err := r.db.Create(application).Error
	return err
}
