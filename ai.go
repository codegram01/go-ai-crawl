package main

import (
	"context"
	"encoding/json"
	"github.com/ollama/ollama/api"
)

func AskAI(mess string, format json.RawMessage) (string, error) {
	var respContent string
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	var stream = false
	req := &api.ChatRequest{
		Model: "llama3.2",
		Messages: []api.Message{
			{
				Role:    "user",
				Content: mess,
			},
		},
		Stream: &stream,
		Format: format,
	}

	respFunc := func(resp api.ChatResponse) error {
		respContent = resp.Message.Content
		return nil
	}
	err = client.Chat(ctx, req, respFunc)

	return respContent, err
}
