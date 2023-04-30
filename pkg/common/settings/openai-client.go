package settings

import (
	openai "github.com/sashabaranov/go-openai"
)

type OpenAiClient struct {
	Client *openai.Client
}

var Client *OpenAiClient

func SetOpenAiClient(key string) {
	Client = &OpenAiClient{Client: openai.NewClient(key)}
}
