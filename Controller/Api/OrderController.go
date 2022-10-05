package Api

import (
	"fmt"
	"ginApi/Common/Tools"
	"ginApi/Controller"
	"ginApi/Service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
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
	if Tools.Config.GetString("token.type") == "token" {
		users, _ := c.Get("users")
		usersMap := users.(map[string]string)
		if userId, ok := usersMap["userId"]; ok {
			param.UserId, _ = strconv.Atoi(usersMap["userId"])
			fmt.Println("%T", userId)
		}
	}
	if Tools.Config.GetString("token.type") == "jwt" {
		param.UserId = c.GetInt("userId")
	}
	data, lastPage, total, _ := Service.OrderService{}.Lists(&param)
	this.Success(c, map[string]interface{}{
		"lists":    data,
		"lastPage": lastPage,
		"total":    total,
	})
}
