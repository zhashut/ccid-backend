package routers

import (
	"github.com/gin-gonic/gin"
	"go-init/api"
)

/**
 * @author: zhashut
 * Date: 2025/4/19
 * Time: 9:53
 * Description: 团队成员-路由层
 */

func InitTeamMemberRouter(r *gin.RouterGroup) {
	g := r.Group("/teamMember")
	{
		g.POST("/create", api.CreateTeamMember)
		g.POST("/update", api.UpdateTeamMember)
		g.POST("/updateRole", api.UpdateMemberRole)
		g.POST("/delete", api.DeleteTeamMember)
		g.GET("/list", api.ListTeamMember)
	}
}
