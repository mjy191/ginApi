package Routers

import (
	"ginApi/Controller/Api"
	"ginApi/Middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (this ApiRouter) Router(r *gin.Engine) {
	apiGroup := r.Group("/api", Middleware.CheckSignMiddleware{}.Handle(), Middleware.CheckTokenMiddleware{}.Handle(), Middleware.CheckJwtMiddleware{}.Handle())
	{
		apiGroup.Any("/user/lists", Api.UserController{}.Lists)
		apiGroup.Any("/user/edit", Api.UserController{}.Edit)
		apiGroup.Any("/user/del", Api.UserController{}.Del)

		apiGroup.Any("/order/lists", Api.OrderController{}.Lists)
	}

	signGroup := r.Group("/api", Middleware.CheckSignMiddleware{}.Handle())
	{
		signGroup.Any("/user/login", Api.UserController{}.Login)
		signGroup.Any("/user/add", Api.UserController{}.Add)
	}
}
