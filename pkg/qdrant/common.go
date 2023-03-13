package qdrant

import (
	"bytes"
	"document-ai/pkg/common"
	"io"
	"net/http"
)

func Send(httpMethod string, suffix string, reqBytes []byte) (body []byte, err error) {
	req, err := http.NewRequest(httpMethod, common.GlobalObject.Qdrant.Host+suffix, bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	return
}

type CommonResponse struct {
	Result interface{} `json:"result"`
	Status interface{} `json:"status"`
	Time   float64     `json:"time"`
}
