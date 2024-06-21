package controller

import "github.com/gin-gonic/gin"

const (
	SUCC = "success"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func ReturnSucc(ctx *gin.Context, data any) {
	resp := new(Response)
	resp.Code = 200
	resp.Msg = "sucess"
	resp.Data = data
	ctx.JSON(200, resp)
}

func ReturnError(ctx *gin.Context, code int, msg string) {
	resp := new(Response)
	resp.Code = code
	resp.Msg = msg
	resp.Data = ""
	ctx.JSON(200, resp)
}
