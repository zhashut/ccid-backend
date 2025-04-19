package routers

import (
	"github.com/gin-gonic/gin"
	"go-init/api"
)

/**
 * @author: zhashut
 * Date: 2025/4/19
 * Time: 0:22
 * Description: 数据报告-路由层
 */

func InitDataProperty(r *gin.RouterGroup) {
	g := r.Group("/dataProperty")
	{
		g.POST("/create", api.CreateDataProperty)
		g.POST("/update", api.UpdateDataProperty)
		g.POST("/delete", api.DeleteDataProperty)
		g.GET("/list", api.ListDataProperty)
		g.GET("/patentList", api.PatentList)
	}
}
