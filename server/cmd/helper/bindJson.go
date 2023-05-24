package helper

import (
	"log"
	"server/cmd/web"

	"github.com/gin-gonic/gin"
)

func BindJSON(ctx *gin.Context, request web.MoviePayloadResponse)  {
	if err := ctx.BindJSON(&request); err != nil {
		log.Println(err)
		return
	}
}