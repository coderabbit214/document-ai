package openai

import (
	"bytes"
	"document-ai/pkg/common"
	"io"
	"net/http"
)

const (
	Api = "https://api.openai.com/v1"
)

// Usage Represents the total token usage per request to OpenAI.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func Send(suffix string, reqBytes []byte) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodPost, Api+suffix, bytes.NewBuffer(reqBytes))
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+common.GlobalObject.OpenAi.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	return
}
