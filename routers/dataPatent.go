package routers

import (
	"github.com/gin-gonic/gin"
	"go-init/api"
)

/**
 * @author: zhashut
 * Date: 2025/4/19
 * Time: 10:55
 * Description: 数据专利-路由层
 */

func InitDataPatentRouter(r *gin.RouterGroup) {
	g := r.Group("/dataPatent")
	{
		g.POST("/create", api.CreateDataPatent)
		g.POST("/update", api.UpdateDataPatent)
		g.POST("/delete", api.DeleteDataPatent)
		g.GET("/list", api.ListDataPatent)
	}
}
