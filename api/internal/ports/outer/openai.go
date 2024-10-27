package outer

type OpenAIService interface {
	GenerateText() (string, error)
}
