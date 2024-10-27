package usecase

import (
	"github.com/ei-sugimoto/yamicheck/api/internal/domain"
)

type PreInspectService interface {
	PreInspect(*domain.Job) (*domain.Level, error)
}

type PreInspectServiceImp struct {
}

func NewInspectService() PreInspectService {
	return &PreInspectServiceImp{}
}

func (s *PreInspectServiceImp) PreInspect(job *domain.Job) (*domain.Level, error) {
	currentLevel := domain.Level(checkHourlyWage(job))

	return &currentLevel, nil

}

func checkHourlyWage(job *domain.Job) int {
	if job.HourlyWage <= 2500 {
		return 1
	} else if job.HourlyWage <= 5000 {
		return 2
	}

	return 3

}
