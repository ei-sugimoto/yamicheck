package domain

import "fmt"

type Level int

func (l Level) string() (string, error) {
	switch l {
	case 1:
		return "safe", nil
	case 2:
		return "warning", nil
	case 3:
		return "danger", nil
	case 4:
		return "critical", nil
	default:
		return "", fmt.Errorf("invalid level: %d", l)
	}
}

func (l *Level) Validate() error {
	if *l < 1 || *l > 4 {
		return fmt.Errorf("invalid level: %d", *l)
	}
	return nil
}

func (l *Level) Integer() int {
	return int(*l)
}
