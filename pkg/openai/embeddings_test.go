package openai

import (
	"reflect"
	"testing"
)

func TestSendEmbeddings(t *testing.T) {
	type args struct {
		request EmbeddingRequest
	}
	tests := []struct {
		name                  string
		args                  args
		wantEmbeddingResponse EmbeddingResponse
		wantErr               bool
	}{
		{
			name: "1",
			args: args{EmbeddingRequest{
				Input: "哈哈哈哈",
				Model: TextEmbeddingAda002,
			}},
			wantEmbeddingResponse: EmbeddingResponse{},
			wantErr:               false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEmbeddingResponse, err := SendEmbeddings(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEmbeddings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEmbeddingResponse, tt.wantEmbeddingResponse) {
				t.Errorf("CreateEmbeddings() gotEmbeddingResponse = %v, want %v", gotEmbeddingResponse, tt.wantEmbeddingResponse)
			}
		})
	}
}
