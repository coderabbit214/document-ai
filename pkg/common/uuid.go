package common

import (
	"github.com/google/uuid"
	"strings"
)

// GenerateUUID 生成uuid
func GenerateUUID() string {
	oldId := uuid.New()
	key := oldId.String()
	id := strings.ReplaceAll(key, "-", "")
	return id
}
