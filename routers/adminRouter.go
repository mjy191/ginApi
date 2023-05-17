package routers

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (this AdminRouter) Router(r *gin.Engine) {
	r.GET("/admin/index/index")
}
