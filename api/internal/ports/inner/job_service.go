package inner

import "github.com/ei-sugimoto/yamicheck/api/internal/domain"

type JobService interface {
	Check(*domain.Job) (*domain.Level, error)
}
