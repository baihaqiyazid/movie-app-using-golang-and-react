package helper

import (
	"net/http"
	"server/cmd/web"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   data,
	})
}