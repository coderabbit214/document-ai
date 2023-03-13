package qdrant

import (
	"document-ai/pkg/common"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	collectionApi = "/collections"
)

type CreateCollectionRequest struct {
	Vectors Vectors `json:"vectors"`
}
type Vectors struct {
	Size     int    `json:"size"`
	Distance string `json:"distance"`
}

func init() {
	err := GetCollection(common.GlobalObject.Qdrant.CollectionName)
	if err == nil {
		return
	}
	createCollectionRequest := CreateCollectionRequest{Vectors: Vectors{
		Size:     1536,
		Distance: "Cosine",
	}}
	err = CreateCollection(common.GlobalObject.Qdrant.CollectionName, createCollectionRequest)
	if err != nil {
		panic("init error:" + err.Error())
	}
}

// CreateCollection creates a collection in qdrant
func CreateCollection(name string, createCollectionRequest CreateCollectionRequest) (err error) {
	response := &CommonResponse{}
	var reqBytes []byte
	reqBytes, err = json.Marshal(createCollectionRequest)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	body, err := Send(http.MethodPut, collectionApi+"/"+name, reqBytes)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}
	if response.Result == nil {
		return errors.New(response.Status.(map[string]interface{})["error"].(string))
	}
	return

}

func GetCollection(name string) (err error) {
	response := &CommonResponse{}

	body, err := Send(http.MethodGet, collectionApi+"/"+name, nil)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}
	if response.Result == nil {
		return errors.New(response.Status.(map[string]interface{})["error"].(string))
	}
	return

}
