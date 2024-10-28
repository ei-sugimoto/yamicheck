package cmd

import (
	"fmt"
	"os"

	"github.com/ei-sugimoto/yamicheck/api/internal/domain"
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
	testJob := domain.Job{
		CompanyName: "テスト株式会社",
		HourlyWage:  3000,
		Description: "テスト株式会社では、誰でも稼げるお仕事を紹介しています。お気軽にお問い合わせください。#UD",
	}
	preLevel, err := usecase.NewInspectService().PreInspect(&testJob)
	if err != nil {
		panic(err)
	}

	openaiService := usecase.NewOpenAIService(client)
	fmt.Println("Generated text:")
	level, err := openaiService.Inspect(&testJob, preLevel)
	if err != nil {
		panic(err)
	}
	fmt.Println(level.Integer())
}
