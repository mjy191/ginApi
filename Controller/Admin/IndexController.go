package Admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

//前后端分离已经不用这种模式
func (this IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "Admin/index.html", gin.H{
		"title": "后台首页",
	})
}
