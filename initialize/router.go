package initialize

import (
	"github.com/gin-gonic/gin"
	"go-init/middlewares"
	"go-init/routers"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/6/7
 * Time: 9:10
 * Description: 统一全局路由初始化
 */

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	api := r.Group("/api")
	routers.InitDataProperty(api)
	routers.InitDataTask(api)
	routers.InitDataReport(api)
	routers.InitDataPatentRouter(api)
	routers.InitTeamMemberRouter(api)
	return r
}
