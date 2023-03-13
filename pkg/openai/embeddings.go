package openai

import (
	"document-ai/pkg/common"
	"encoding/json"
)

const (
	embeddingsApi = "/embeddings"
)
const (
	TextEmbeddingAda002 = "text-embedding-ada-002"
)

type EmbeddingRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

// EmbeddingResponse is the response from a Create embeddings request.
type EmbeddingResponse struct {
	Object string      `json:"object"`
	Data   []Embedding `json:"data"`
	Model  string      `json:"model"`
	Usage  Usage       `json:"usage"`
}

type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

// SendEmbeddings returns an EmbeddingResponse
func SendEmbeddings(request EmbeddingRequest) (embeddingResponse EmbeddingResponse, err error) {
	var reqBytes []byte
	reqBytes, err = json.Marshal(request)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	body, err := Send(embeddingsApi, reqBytes)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	err = json.Unmarshal(body, &embeddingResponse)
	return
}
