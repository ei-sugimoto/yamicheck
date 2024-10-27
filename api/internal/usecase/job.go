package usecase

import (
	"github.com/ei-sugimoto/yamicheck/api/internal/domain"
	"github.com/ei-sugimoto/yamicheck/api/internal/ports/inner"
)

type JobServiceImp struct {
}

func NewJobService() inner.JobService {
	return &JobServiceImp{}
}

func (s *JobServiceImp) Check(job *domain.Job) (*domain.Level, error) {
	preCheckService := NewInspectService()
	currentLevel, err := preCheckService.PreInspect(job)
	if err != nil {
		return nil, err
	}

	return currentLevel, nil
}
