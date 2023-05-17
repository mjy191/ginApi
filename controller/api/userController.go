package api

import (
	"ginApi/common/config"
	"ginApi/common/enum"
	"ginApi/common/jwt"
	"ginApi/common/response"
	"ginApi/common/tools"
	"ginApi/controller"
	"ginApi/models"
	"ginApi/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

type UserController struct {
	controller.BaseController
}

func (this UserController) Lists(c *gin.Context) {
	var userParam service.AddParam
	err := c.ShouldBindBodyWith(&userParam, binding.JSON)
	if err != nil {
		panic(err)
	}
	data, _ := service.UserService{}.Lists(&userParam)
	response.Success(c, &response.Response{Data: data})
}

func (this UserController) Add(c *gin.Context) {
	var param service.AddParam
	if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		tools.GetError(err, param)
		return
	}
	id := service.UserService{}.Add(&param)
	response.Success(c, &response.Response{
		Data: map[string]int{"id": id},
	})
}

func (this UserController) Edit(c *gin.Context) {
	var param service.EditParam
	if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		tools.GetError(err, param)
		return
	}
	service.UserService{}.Edit(&param)

	response.Success(c, &response.Response{
		Data: map[string]string{},
	})
}

func (this UserController) Del(c *gin.Context) {
	var delParam service.DelParam
	err := c.ShouldBindBodyWith(&delParam, binding.JSON)
	if err != nil {
		panic(err)
	}
	service.UserService{}.Del(&delParam)

	response.Success(c, &response.Response{
		Data: map[string]string{},
	})
}

func (this UserController) Login(c *gin.Context) {
	var loginParam service.LoginParam
	err := c.ShouldBindBodyWith(&loginParam, binding.JSON)
	if err != nil {
		tools.GetError(err, loginParam)
	}
	data, _ := service.UserService{}.Login(&loginParam)

	// 生成token
	var token string
	if config.Viper.GetString("token.type") == "jwt" {
		result, err := jwt.Jwt{}.CreateToken(data.Id)
		if err != nil {
			response.Fail(c, &response.Response{Code: enum.CodeParamError, Msg: enum.ErrMsg[enum.CodeSystemError]})
			return
		}
		token = result
	} else if config.Viper.GetString("token.type") == "token" {
		var res map[string]string
		for {
			token = tools.RandString(config.Viper.GetInt("token.length"))
			res, _ = models.RedisDb.HGetAll(token).Result()
			if len(res) != 0 {
				token = tools.RandString(config.Viper.GetInt("token.length"))
			} else {
				break
			}
		}
		models.RedisDb.HMSet("token:"+token, map[string]interface{}{
			"userId": data.Id,
			"token":  token,
		})

		models.RedisDb.Expire(
			"token:"+token,
			time.Duration(config.Viper.GetInt64("token.expire"))*time.Second,
		)
	} else {
		response.Fail(c, &response.Response{Code: enum.CodeSystemError, Msg: "token参数配置错误"})
	}

	response.Success(c, &response.Response{
		Data: map[string]string{"token": token},
	})
}
