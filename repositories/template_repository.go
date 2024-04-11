package repositories

import (
	"github.com/ribbinpo/mining-mine-bot/domain"
	"gorm.io/gorm"
)

type templateRepo struct {
	Db *gorm.DB
}

func NewTemplateRepository(db *gorm.DB) domain.TemplateRepository {
	return &templateRepo{Db: db}
}

func (t *templateRepo) GetAll() ([]domain.Template, error) {
	var templates []domain.Template
	err := t.Db.Find(&templates).Error
	if err != nil {
		return nil, err
	}
	return templates, nil
}
