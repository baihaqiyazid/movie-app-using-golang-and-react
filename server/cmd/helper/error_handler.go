package helper

import (
	"net/http"
	"server/cmd/web"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx gin.Context, err error)  {
	
	ctx.JSON(http.StatusBadRequest, web.WebResponse{
		Code: http.StatusBadRequest,
		Status: "BAD REQUEST",
	})
}

func NotFound(ctx gin.Context, err error)  {
	
	ctx.JSON(http.StatusNotFound, web.WebResponse{
		Code: http.StatusNotFound,
		Status: "DATA NOT FOUND",
	})
}

func PanicError(err error)  {
	if err != nil {
		panic(err)
	}
}