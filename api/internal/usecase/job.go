package usecase

import (
	"os"

	"github.com/ei-sugimoto/yamicheck/api/internal/domain"
	"github.com/ei-sugimoto/yamicheck/api/internal/ports/inner"
	"github.com/sashabaranov/go-openai"
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

	if currentLevel.Integer() == 4 {
		return currentLevel, nil
	}

	token := os.Getenv("OPENAI_API_KEY")

	if token == "" {
		panic("OPENAI_API_KEY is not set")
	}

	client := openai.NewClient(token)
	openaiService := NewOpenAIService(client)
	level, err := openaiService.Inspect(job, currentLevel)
	if err != nil {
		return nil, err
	}

	return level, nil
}
