package outer

import "github.com/ei-sugimoto/yamicheck/api/internal/domain"

type OpenAIService interface {
	GenerateText() (string, error)
	Inspect(*domain.Job, *domain.Level) (*domain.Level, error)
}
