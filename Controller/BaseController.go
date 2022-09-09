package Controller

import (
	"ginApi/Common/Enum"
	"ginApi/Common/Logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (this BaseController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":  Enum.CodeSuccess,
		"msg":   Enum.ErrMsg[Enum.CodeSuccess],
		"logid": Logger.Logid,
		"data":  data,
	})
}

func (this BaseController) Fail(c *gin.Context, data map[string]interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":  data["code"],
		"logid": Logger.Logid,
		"msg":   data["msg"],
	})
}
