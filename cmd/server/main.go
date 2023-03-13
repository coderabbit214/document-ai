package main

import (
	"document-ai/pkg/api"
	"document-ai/pkg/common"
	"fmt"
)

func main() {
	r := api.SetupRouter()
	if err := r.Run(":" + common.GlobalObject.Server.Port); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
