package api

import (
	"document-ai/pkg/common"
	"document-ai/pkg/openai"
	"document-ai/pkg/qdrant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func updateByJson(c *gin.Context) {
	var json []map[string]string
	if err := c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(json) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json is empty"})
		return
	}
	//数据向量化
	points := make([]qdrant.Point, 0)
	for _, v := range json {
		embeddingRequest := openai.EmbeddingRequest{
			Input: v["text"],
			Model: openai.TextEmbeddingAda002,
		}
		response, err := openai.SendEmbeddings(embeddingRequest)
		if err != nil {
			common.Logger.Error(err.Error())
			c.JSON(http.StatusOK, common.Error(err.Error()))
			return
		}
		points = append(points, qdrant.Point{
			ID:      uuid.New().String(),
			Payload: v,
			Vector:  response.Data[0].Embedding,
		})
	}
	pr := qdrant.PointRequest{
		Points: points,
	}

	//存储
	err := qdrant.CreatePoints(common.GlobalObject.Qdrant.CollectionName, pr)
	if err != nil {
		common.Logger.Error(err.Error())
		c.JSON(http.StatusOK, common.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.Success(nil))
}
