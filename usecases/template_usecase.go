package usecases

import (
	"github.com/ribbinpo/mining-mine-bot/domain"
)

// Service
type TemplateService struct {
	repo domain.TemplateRepository
}

// Constructor
func NewTemplateService(repo domain.TemplateRepository) domain.TemplateUseCase {
	return &TemplateService{repo: repo}
}

func (s *TemplateService) GetAll() ([]domain.Template, error) {
	return s.repo.GetAll()
}
