package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ei-sugimoto/yamicheck/api/internal/domain"
	"github.com/ei-sugimoto/yamicheck/api/internal/ports/outer"
	"github.com/sashabaranov/go-openai"
)

type OpenAIServiceImp struct {
	Client *openai.Client
}

func NewOpenAIService(cilent *openai.Client) outer.OpenAIService {
	return &OpenAIServiceImp{
		Client: cilent,
	}
}

func (s *OpenAIServiceImp) GenerateText() (string, error) {
	resp, err := s.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "こんにちは",
			},
		},
	})

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (s *OpenAIServiceImp) Inspect(job *domain.Job, preLevel *domain.Level) (*domain.Level, error) {
	headerStr := "今から、求人情報を共有します。あなたは、この求人情報を参照して犯罪に関わる可能性がある求人かどうかを判断します。判断する基準として、4段階のレベルを設定しています。レベル1は安全、レベル2は警告、レベル3は危険、レベル4は重大です。このチャットの返答については、このチャットの返答については、レベルの数値のみを返答してください。(例:3)\n"
	subHeaderStr := fmt.Sprintf("また求人情報には、事前に精査しており、その結果はすでにレベル%dとして設定されています。なので、このレベル未満の返答はしないようにしてください。\n", preLevel.Integer())

	dangerKeywords := []string{"合法", "たたき", "叩き", "タヌキ", "炊飯器", "かけ子", "掛け", "受け子", "#UD", "出し子"}
	KeywordsStr := fmt.Sprintf("また危険なキーワードとして次に列挙する言葉が入っている場合には無条件でレベルを1段階上げてください。また二つ含まれる場合はレベルを2段階上げてください\n 危険なキーワード: %s\n", dangerKeywords)
	contentStr := fmt.Sprintf("求人情報: %s\n", job.Description)
	footerStr := "以上です。よろしくお願いします。"
	content := headerStr + subHeaderStr + KeywordsStr + contentStr + footerStr
	resp, err := s.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	levelStr := resp.Choices[0].Message.Content
	level, err := strconv.Atoi(levelStr)
	if err != nil {
		return nil, err
	}
	newLevel := domain.Level(level)
	if err := newLevel.Validate(); err != nil {
		return nil, err
	}

	return &newLevel, nil
}
