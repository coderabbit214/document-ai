package api

import (
	"document-ai/pkg/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Healthy(c *gin.Context) {
	c.JSON(http.StatusOK, common.Success("healthy!"))
}

func Error(c *gin.Context) {
	common.Logger.Debug("Debug")
	common.Logger.Info("Info")
	common.Logger.Warn("Warn")
	common.Logger.Error("Error")
	c.JSON(http.StatusOK, common.Error("Error"))
}
