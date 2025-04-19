package routers

import (
	"github.com/gin-gonic/gin"
	"go-init/api"
)

/**
 * @author: zhashut
 * Date: 2025/4/19
 * Time: 0:41
 * Description: 数据报告-路由层
 */

func InitDataReport(r *gin.RouterGroup) {
	g := r.Group("/dataReport")
	{
		g.POST("/create", api.CreateDataReport)
		g.POST("/update", api.UpdateDataReport)
		g.POST("/updateStatus", api.UpdateReportStatus)
		g.POST("/delete", api.DeleteDataReport)
		g.GET("/list", api.ListDataReport)
	}
}
