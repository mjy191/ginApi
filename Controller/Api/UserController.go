package Api

import (
	"ginApi/Common/Enum"
	"ginApi/Common/Tools"
	"ginApi/Controller"
	"ginApi/Models"
	"ginApi/Service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

type UserController struct {
	Controller.BaseController
}

func (this UserController) Lists(c *gin.Context) {
	var userParam Service.AddParam
	err := c.ShouldBindBodyWith(&userParam, binding.JSON)
	if err != nil {
		panic(err)
	}
	data, _ := Service.UserService{}.Lists(&userParam)
	this.Success(c, data)
}

func (this UserController) Add(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			this.Fail(c, r.(map[string]interface{}))
			return
		}
	}()
	var param Service.AddParam
	if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		Tools.GetError(err, param)
		return
	}
	id := Service.UserService{}.Add(&param)
	this.Success(c, map[string]int{"id": id})
}

func (this UserController) Edit(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			this.Fail(c, r.(map[string]interface{}))
			return
		}
	}()
	var param Service.EditParam
	if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		Tools.GetError(err, param)
		return
	}
	Service.UserService{}.Edit(&param)
	this.Success(c, map[string]string{})
}

func (this UserController) Del(c *gin.Context) {
	var delParam Service.DelParam
	err := c.ShouldBindBodyWith(&delParam, binding.JSON)
	if err != nil {
		panic(err)
	}
	Service.UserService{}.Del(&delParam)
	this.Success(c, map[string]string{})
}

func (this UserController) Login(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			this.Fail(c, r.(map[string]interface{}))
			return
		}
	}()
	var loginParam Service.LoginParam
	err := c.ShouldBindBodyWith(&loginParam, binding.JSON)
	if err != nil {
		Tools.GetError(err, loginParam)
	}
	data, _ := Service.UserService{}.Login(&loginParam)

	// 生成token
	var token string
	if Tools.Config.GetString("token.type") == "jwt" {
		result, err := Tools.Jwt{}.CreateToken(data.Id)
		if err != nil {
			this.Fail(c, map[string]interface{}{
				"code": Enum.CodeSystemError,
				"msg":  Enum.ErrMsg[Enum.CodeSystemError],
			})
			return
		}
		token = result
	} else if Tools.Config.GetString("token.type") == "token" {
		var res map[string]string
		for {
			token = Tools.RandString(Tools.Config.GetInt("token.length"))
			res, _ = Models.RedisDb.HGetAll(token).Result()
			if len(res) != 0 {
				token = Tools.RandString(Tools.Config.GetInt("token.length"))
			} else {
				break
			}
		}
		Models.RedisDb.HMSet("token:"+token, map[string]interface{}{
			"userId": data.Id,
			"token":  token,
		})

		Models.RedisDb.Expire(
			"token:"+token,
			time.Duration(Tools.Config.GetInt64("token.expire"))*time.Second,
		)
	} else {
		this.Fail(c, map[string]interface{}{
			"code": Enum.CodeSystemError,
			"msg":  "token参数配置错误",
		})
	}

	this.Success(c, map[string]string{
		"token": token,
	})
}
