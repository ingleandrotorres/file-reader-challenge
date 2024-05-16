package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ping struct {
}

func NewPingController() Ping {
	return Ping{}
}
func (c Ping) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
