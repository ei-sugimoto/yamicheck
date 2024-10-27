package cmd

import (
	"fmt"
	"os"

	"github.com/ei-sugimoto/yamicheck/api/internal/usecase"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func TmpTask() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	token := os.Getenv("OPENAI_API_KEY")

	if token == "" {
		panic("OPENAI_API_KEY is not set")
	}

	client := openai.NewClient(token)
	openaiService := usecase.NewOpenAIService(client)
	fmt.Println("Generated text:")
	fmt.Println(openaiService.GenerateText())
}
