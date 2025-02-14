package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// ResError Return error information
func ResError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK,
		&Data{
			Code: code,
			Msg:  code.Msg(),
			Data: nil,
		})
}

// ResErrorWithMsg Custom error return
func ResErrorWithMsg(c *gin.Context, code ResCode, msg interface{}, data ...interface{}) {
	c.JSON(http.StatusOK,
		&Data{
			Code: code,
			Msg:  msg,
			Data: data,
		})
}

// ResSuccess Return a success message
func ResSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK,
		&Data{
			Code: CodeSuccess,
			Msg:  CodeSuccess.Msg(),
			Data: data,
		})
}
