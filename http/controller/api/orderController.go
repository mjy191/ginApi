package api

import (
	"ginApi/common/response"
	"ginApi/common/tools"
	"ginApi/http/controller"
	"ginApi/http/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrderController struct {
	controller.BaseController
}

func (this OrderController) Lists(c *gin.Context) {
	var param service.OrderParam
	// 根据 body为空会报错EOF err.Error() != "EOF" ,查询条件是否必填为判断
	if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil && err.Error() != "EOF" {
		tools.GetError(err, param)
		return
	}
	param.UserId = c.GetInt("userId")
	data, lastPage, total, _ := service.OrderService{}.Lists(&param)
	response.Success(c, &response.Response{Data: map[string]interface{}{
		"lists":    data,
		"lastPage": lastPage,
		"total":    total,
	}})
}
