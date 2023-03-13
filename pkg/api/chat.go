package api

import (
	"document-ai/pkg/common"
	"document-ai/pkg/openai"
	"document-ai/pkg/qdrant"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Message struct {
	Id   string `json:"id"`
	Text string `json:"content"`
}

func chat(c *gin.Context) {
	var message Message
	if err := c.Bind(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//计算向量
	embeddingRequest := openai.EmbeddingRequest{
		Input: message.Text,
		Model: openai.TextEmbeddingAda002,
	}
	response, err := openai.SendEmbeddings(embeddingRequest)
	if err != nil {
		common.Logger.Error(err.Error())
		c.JSON(http.StatusOK, common.Error(err.Error()))
		return
	}
	params := make(map[string]interface{})
	params["exact"] = false
	params["hnsw_ef"] = 128

	sr := qdrant.PointSearchRequest{
		Params:      params,
		Vector:      response.Data[0].Embedding,
		Limit:       3,
		WithPayload: true,
	}
	//查询相似的
	res, err := qdrant.SearchPoints(common.GlobalObject.Qdrant.CollectionName, sr)
	if err != nil {
		common.Logger.Error(err.Error())
		c.JSON(http.StatusOK, common.Error(err.Error()))
		return
	}
	//组装本地数据
	localData := ""
	for i, v := range res {
		re := v.Payload.(map[string]interface{})
		localData += "\n"
		localData += strconv.Itoa(i)
		localData += "."
		localData += re["title"].(string)
		localData += ":"
		localData += re["text"].(string)
	}
	messages := make([]openai.ChatCompletionMessage, 0)
	q := "使用以下段落来回答问题，如果段落内容不相关就返回未查到相关信息：\"" + message.Text + "\""
	q += localData

	system := openai.ChatCompletionMessage{
		Role:    "system",
		Content: "你是一个医院问诊机器人",
	}
	demo_q := "使用以下段落来回答问题：\"成人头疼，流鼻涕是感冒还是过敏？\"\n1. 普通感冒：您会出现喉咙发痒或喉咙痛，流鼻涕，流清澈的稀鼻涕（液体），有时轻度发热。\n2. 常年过敏：症状包括鼻塞或流鼻涕，鼻、口或喉咙发痒，眼睛流泪、发红、发痒、肿胀，打喷嚏。"
	demo_a := "成人出现头痛和流鼻涕的症状，可能是由于普通感冒或常年过敏引起的。如果病人出现咽喉痛和咳嗽，感冒的可能性比较大；而如果出现口、喉咙发痒、眼睛肿胀等症状，常年过敏的可能性比较大。"
	user1 := openai.ChatCompletionMessage{
		Role:    "user",
		Content: demo_q,
	}
	assistant := openai.ChatCompletionMessage{
		Role:    "assistant",
		Content: demo_a,
	}
	user := openai.ChatCompletionMessage{
		Role:    "user",
		Content: q,
	}

	messages = append(messages, system)
	messages = append(messages, user1)
	messages = append(messages, assistant)
	messages = append(messages, user)
	chatCompletionRequest := openai.ChatCompletionRequest{
		Model:    openai.Gpt3Dot5Turbo,
		Messages: messages,
	}
	completionResponse, err := openai.SendChat(chatCompletionRequest)
	if err != nil {
		common.Logger.Error(err.Error())
		c.JSON(http.StatusOK, common.Error(err.Error()))
		return
	}
	content := completionResponse.Choices[0].Message.Content
	c.JSON(http.StatusOK, common.Success(content))
	return
}
