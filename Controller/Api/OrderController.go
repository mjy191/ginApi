package Api

import (
	"ginApi/Common/Tools"
	"ginApi/Controller"
	"ginApi/Service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrderController struct {
	Controller.BaseController
}

func (this OrderController) Lists(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			this.Fail(c, r.(map[string]interface{}))
			return
		}
	}()
	var param Service.OrderParam
	if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		Tools.GetError(err, param)
		return
	}
	param.UserId = c.GetInt("userId")
	data, lastPage, total, _ := Service.OrderService{}.Lists(&param)
	this.Success(c, map[string]interface{}{
		"lists":    data,
		"lastPage": lastPage,
		"total":    total,
	})
}
