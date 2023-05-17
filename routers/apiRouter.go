package routers

import (
	"ginApi/controller/api"
	"ginApi/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (this ApiRouter) Router(r *gin.Engine) {
	apiGroup := r.Group("/api", middleware.CheckSignMiddleware{}.Handle(), middleware.CheckTokenMiddleware{}.Handle(), middleware.CheckJwtMiddleware{}.Handle())
	{
		apiGroup.Any("/user/lists", api.UserController{}.Lists)
		apiGroup.Any("/user/edit", api.UserController{}.Edit)
		apiGroup.Any("/user/del", api.UserController{}.Del)

		apiGroup.Any("/order/lists", api.OrderController{}.Lists)
	}

	signGroup := r.Group("/api", middleware.CheckSignMiddleware{}.Handle())
	{
		signGroup.Any("/user/login", api.UserController{}.Login)
		signGroup.Any("/user/add", api.UserController{}.Add)
	}
}
