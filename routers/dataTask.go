package routers

import (
	"github.com/gin-gonic/gin"
	"go-init/api"
)

/**
 * @author: zhashut
 * Date: 2025/4/19
 * Time: 0:39
 * Description: 数据任务-路由层
 */

func InitDataTask(r *gin.RouterGroup) {
	g := r.Group("/dataTask")
	{
		g.POST("/create", api.CreateDataTask)
		g.POST("/update", api.UpdateDataTask)
		g.POST("/updateStatus", api.UpdateTaskStatus)
		g.POST("/delete", api.DeleteDataTask)
		g.GET("/list", api.ListDataTask)
	}
}
