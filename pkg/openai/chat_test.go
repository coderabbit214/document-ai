package openai

import (
	"reflect"
	"testing"
)

func TestSendChat(t *testing.T) {
	type args struct {
		request ChatCompletionRequest
	}
	var tests = []struct {
		name                       string
		args                       args
		wantChatCompletionResponse ChatCompletionResponse
		wantErr                    bool
	}{
		{
			name: "1",
			args: args{ChatCompletionRequest{
				MaxTokens: 5,
				Model:     Gpt3Dot5Turbo,
				Messages: []ChatCompletionMessage{
					{
						Role:    ChatMessageRoleUser,
						Content: "Hello!",
					},
				},
			}},
			wantChatCompletionResponse: ChatCompletionResponse{},
			wantErr:                    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChatCompletionResponse, err := SendChat(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotChatCompletionResponse, tt.wantChatCompletionResponse) {
				t.Errorf("SendChat() gotChatCompletionResponse = %v, want %v", gotChatCompletionResponse, tt.wantChatCompletionResponse)
			}
		})
	}
}
